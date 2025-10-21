package kafka

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/confluentinc/confluent-kafka-go/kafka"
// )

// func produceMessage(broker, topic, message string) error {
// 	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer p.Close()

// 	// Delivery report handler (runs in a goroutine)
// 	go func() {
// 		for e := range p.Events() {
// 			switch ev := e.(type) {
// 			case *kafka.Message:
// 				if ev.TopicPartition.Error != nil {
// 					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
// 				} else {
// 					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
// 				}
// 			}
// 		}
// 	}()

// 	// Send a message
// 	err = p.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          []byte("Hello Kafka from Go!"),
// 	}, nil)

// 	if err != nil {
// 		fmt.Printf("Produce error: %v\n", err)
// 	}

// 	// Wait for message deliveries before shutting down
// 	p.Flush(5000)
// }

// func consumeMessages(broker, topic string) error {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": broker,
// 		"group.id":          "go-kafka-group",
// 		"auto.offset.reset": "earliest",
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Close()

// 	err = c.SubscribeTopics([]string{topic}, nil)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Waiting for messages... (Ctrl+C to quit)")

// 	sigchan := make(chan os.Signal, 1)
// 	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

// runLoop:
// 	for {
// 		select {
// 		case sig := <-sigchan:
// 			fmt.Printf("Caught signal %v: terminating\n", sig)
// 			break runLoop
// 		default:
// 			msg, err := c.ReadMessage(100 * time.Millisecond)
// 			if err == nil {
// 				fmt.Printf("Received message: %s from %s\n", string(msg.Value), msg.TopicPartition)
// 				break runLoop // Stop after one message
// 			} else if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() != kafka.ErrTimedOut {
// 				fmt.Printf("Consumer error: %v\n", err)
// 			}
// 		}
// 	}

// 	return nil
// }
