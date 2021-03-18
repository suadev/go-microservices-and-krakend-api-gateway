package main

import (
	"github.com/suadev/microservices/services.notification/src/kafka"
	"github.com/suadev/microservices/shared/config"
)

func main() {
	config := config.LoadConfig(".")
	kafka.RegisterConsumer(config.KafkaOrderTopic)
}
