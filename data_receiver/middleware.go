package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/toll-calculator/types"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{next: next}
}

func (lm *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obu_id": data.OBUID,
			"lat":    data.Lat,
			"long":   data.Long,
			"took":   time.Since(start),
		}).Info("Producing to Kafka")
	}(time.Now())
	return lm.next.ProduceData(data)

}
