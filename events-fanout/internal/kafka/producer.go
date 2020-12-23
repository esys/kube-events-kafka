package kafka

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Producer interface {
	Write(topic string, msg []byte) error
	Close()
}

type producer struct {
	endpoint string
	*kafka.Producer
}

func NewProducer(endpoint string) (Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": endpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot initialize Kafka producer with endpoint %s: %w", endpoint, err)
	}
	instance := &producer{endpoint, p}

	go func() {
		for e := range instance.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					zap.S().Errorf("delivery failed: %v\n", ev.TopicPartition.Error)
				}
			}
		}
		zap.S().Info("exiting producer events loop")
	}()
	return instance, nil
}

func (p producer) Write(topic string, msg []byte) error {
	km := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}
	if err := p.Produce(&km, nil); err != nil {
		return fmt.Errorf("cannot write message to topic: %w", err)
	}
	return nil
}

func (p producer) Close() {
	zap.S().Info("producer will close")
	p.Producer.Close()
	zap.S().Info("producer closed")
}
