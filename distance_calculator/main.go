package main

import "log"

// type DistanceCalculator struct{
// 	consumer DataConsumer
// }

const kafkaTopic = "obudata"

func main() {
	srv := NewCalculatorService()

	srv = NewLogMiddleware(srv)

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, srv)

	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()

}
