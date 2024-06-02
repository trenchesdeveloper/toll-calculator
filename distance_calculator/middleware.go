package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/toll-calculator/types"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{next: next}
}

func (lm *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"lat":  data.Lat,
			"long": data.Long,
			"err":  err,
		}).Info("Distance calculated")
	}(time.Now())

	return lm.next.CalculateDistance(data)
}
