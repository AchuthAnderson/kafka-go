package config

import (
"github.com/confluentinc/confluent-kafka-go/kafka"
)


const (
	KafkaServer = "localhost:9092"
	KafkaTopic = "orders-v1-topic"
)

config := kafka.ConfigMap