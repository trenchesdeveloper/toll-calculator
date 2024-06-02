package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/toll-calculator/types"
)

// This can also be called KafkaTransport
type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string, svc CalculatorServicer) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
	}, nil

}

func (kc *KafkaConsumer) readMessageLoop() {
	kc.isRunning = true
	for kc.isRunning {
		msg, err := kc.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("Messa  ge on %s", err)
			continue
		}

		var data types.OBUData

		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error: %s", err)
			continue
		}

		distance, err := kc.calcService.CalculateDistance(data)

		if err != nil {
			logrus.Errorf("Error calculating distance: %s", err)
			continue
		}

		_ = distance
	}
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka transport started")
	c.isRunning = true
	c.readMessageLoop()
}
