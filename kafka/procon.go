package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func ProduceAndConsumeMessages() error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	})

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello from Go using apache kafka-go!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("✅ Message sent successfully")
	writer.Close()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "go-consumer-group",
	})

	defer reader.Close()

	fmt.Println("🚀 Listening for messages...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}
}

func main() {
	ProduceAndConsumeMessages()
}
