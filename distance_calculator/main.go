package main

import "log"

// type DistanceCalculator struct{
// 	consumer DataConsumer
// }

const kafkaTopic = "obudata"

func main() {
	srv := NewCalculatorService()
	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, srv)

	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()

}
