package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "test-group",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Error creating consumer: %s\n", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	consumer.Subscribe("test", nil)
	for {
		select {
		case <-sigChan:
			log.Println("Shutting down consumer")
			consumer.Close()
			os.Exit(0)
		default:
			event := consumer.Poll(100)
			if event == nil {
				continue
			}

			switch e := event.(type) {
			case *kafka.Message:
				log.Printf("Received message with value: %s\n", e.Value)
			case kafka.OffsetsCommitted:
				log.Printf("Offsets committed: %s\n", e)
			}
		}
	}
}
