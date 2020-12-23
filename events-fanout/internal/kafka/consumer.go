package kafka

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Consumer interface {
	Read() ([]byte, error)
	Close()
}

type consumer struct {
	endpoint string
	topic string
	count int
	*kafka.Consumer
}

func NewConsumer(endpoint string, topic string) (Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": endpoint,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("cannot initialize Kafka consumer with endpoint %s: %w", endpoint, err)
	}
	if err := c.Subscribe(topic, nil); err != nil {
		return nil, fmt.Errorf("cannot subscribe to topic %s: %w", topic, err)
	}
	return &consumer{endpoint, topic, 0, c}, nil
}

func (c consumer) Read() ([]byte, error) {
	km, err := c.ReadMessage(15 * 1000)
	if err != nil {
		return nil, err
	}
	return km.Value, nil
}

func (c consumer) Close() {
	zap.S().Info("consumer will close")
	if err := c.Consumer.Close(); err != nil {
		zap.S().Errorf("fail to close consumer: %v", err)
	}
	zap.S().Info("consumer closed")
}
