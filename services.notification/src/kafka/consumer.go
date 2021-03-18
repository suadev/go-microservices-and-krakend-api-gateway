package kafka

import (
	"log"
	"os"
	"os/signal"

	"github.com/suadev/microservices/shared/kafka"
)

func RegisterConsumer(topic string) {

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
				if eventType == "OrderCompleted" {
					log.Println("[Notification]: Order Completed!")
				} else if eventType == "OrderFailed" {
					log.Println("[Notification]: Order Failed!")
				}
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
}
