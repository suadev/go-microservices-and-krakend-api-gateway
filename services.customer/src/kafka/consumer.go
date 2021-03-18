package kafka

import (
	"log"
	"os"
	"os/signal"

	eventhandler "github.com/suadev/microservices/services.customer/src/event_handler"
	customer "github.com/suadev/microservices/services.customer/src/internal"
	"github.com/suadev/microservices/shared/kafka"
)

func RegisterConsumer(topic string, service *customer.Service) {

	partitionConsumer, _ := kafka.CreatePartitionConsumer(topic)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				log.Println(err)
			case msg := <-partitionConsumer.Messages():
				log.Println("Message Received:", string(msg.Key), string(msg.Value))
				eventType := string(msg.Headers[0].Value)

				if eventType == "UserCreated" {
					eventhandler.CreateCustomer(service, msg.Value)
				} else if eventType == "ProductCreated" {
					eventhandler.CreateProduct(service, msg.Value)
				} else if eventType == "OrderCompleted" {
					eventhandler.ClearBasket(service, msg.Value)
				}
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
}
