package kafka

import (
	"log"
	"os"
	"os/signal"

	eventhandler "github.com/suadev/microservices/services.product/src/event_handler"
	"github.com/suadev/microservices/services.product/src/internal"
	"github.com/suadev/microservices/shared/kafka"
)

func RegisterConsumer(topic string, service *product.Service) {

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
				if eventType == "OrderCreated" {
					eventhandler.ReserveProducts(service, msg.Value)
				}
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
}
