package main

import (
	"fmt"
	"kafka-go/dto"
	"kafka-go/config"
)

/*
	TODO:
		Producer
			- Write a Simple message to a Topic
		Consumer
			- Consumer Group
				- Start Reading from Multiple Partitions.
			- Read a Simple message from Topic
		Topic:
			- Creating new Topics
*/
func main() {
	fmt.Println("Learning Kafka with GO")
	order := dto.Order{}
	fmt.Println(order)
	fmt.Printf("Bootstrap server is {} and topic being used is {}",config.KafkaServer, config.KafkaTopic)
}