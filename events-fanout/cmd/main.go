package main

import (
	"events-fanout/internal/event"
	"events-fanout/internal/kafka"
	"flag"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Configuration struct {
	Endpoint string `required:"true"`
	Topic    string `required:"true"`
	GroupID  string `default:"events-fanout"`
}

func configure(c *Configuration) {
	if err := envconfig.Process("", c); err != nil {
		zap.S().Fatal(err.Error())
	}
	zap.S().Infof("Endpoint: %s, Topic: %s", c.Endpoint, c.Topic)
}

func setupLogger() {
	logLevel := zap.LevelFlag("logLevel", zap.InfoLevel, "log level: all, debug, info, warn, error, panic, fatal, none")
	flag.Parse()
	zc := zap.NewProductionConfig()
	zc.Level = zap.NewAtomicLevelAt(*logLevel)

	logger, _ := zc.Build()
	zap.ReplaceGlobals(logger)
}

func main() {
	setupLogger()
	sugar := zap.S()

	var cfg Configuration
	configure(&cfg)

	p, err := kafka.NewProducer(cfg.Endpoint)
	if err != nil {
		sugar.Fatal("cannot create producer")
	}
	defer p.Close()

	c, err := kafka.NewConsumer(cfg.Endpoint, cfg.Topic)
	if err != nil {
		sugar.Fatal("cannot create consumer")
	}
	defer c.Close()

	run := true
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		sugar.Infof("signal %s received, terminating", sig)
		run = false
	}()

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for run {
			data, err := c.Read()
			if err != nil {
				sugar.Errorf("read event error: %v", err)
				time.Sleep(5 * time.Second)
				continue
			}
			if data == nil {
				continue
			}
			msg, err := event.CreateDestinationMessage(data)
			if err != nil {
				sugar.Errorf("cannot create destination event: %v", err)
			}
			p.Write(msg.Topic, msg.Message)
		}
		sugar.Info("worker thread done")
		wg.Done()
	}()

	wg.Wait()
}
