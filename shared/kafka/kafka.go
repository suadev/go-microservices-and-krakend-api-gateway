package kafka

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/suadev/microservices/shared/config"
)

func Publish(key uuid.UUID, payload []byte, eventType string, topic string) {

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 3
	saramaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{config.AppConfig.KafkaBrokerAddress}, saramaConfig)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(payload),
		Key:   sarama.StringEncoder(key.String()),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("EventType"),
				Value: []byte(eventType),
			}},
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("[%s] message is stored in topic(%s)/partition(%d)/offset(%d)\n", eventType, topic, partition, offset)
}

func CreatePartitionConsumer(topic string) (sarama.PartitionConsumer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{config.AppConfig.KafkaBrokerAddress}, saramaConfig)
	if err != nil {
		log.Panic(err)
	}
	return consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
}
