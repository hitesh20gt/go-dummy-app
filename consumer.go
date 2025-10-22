package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consume() {

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
