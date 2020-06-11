package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	config := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092"}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Error creating producer: %s\n", err)
	}

	topic := "test"
	record := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("hello world!"),
	}

	producer.ProduceChannel() <- record
	defer producer.Close()

	event := <-producer.Events()
	msg := event.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		log.Printf("Error sending message to cluster: %s\n", msg.TopicPartition.Error)
	} else {
		log.Printf("Message sent to topic %s (partition %d) at offset %d\n",
			*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	}
}
