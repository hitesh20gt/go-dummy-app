package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func Produce() {
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
}
