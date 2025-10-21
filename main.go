package main

import (
	controllers "go-clean-app/controller"
	mongoConnection "go-clean-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoConnection.ConnectMongo()
	r := gin.Default()

	// broker := "localhost:9092"
	// topic := "PLAINTEXT"

	// Produce a message
	// err := kafkaConnection.produceMessage(broker, topic, "Hello from Go!")
	// if err != nil {
	// 	fmt.Printf("Producer error: %v\n", err)
	// 	return
	// }

	// // Consume the message
	// err = kafkaConnection(broker, topic)
	// if err != nil {
	// 	fmt.Printf("Consumer error: %v\n", err)
	// }

	// Route to get dummy user
	r.GET("/user", controllers.CreateAndStoreDummyUser)
	r.GET("/users", controllers.GetAllUsers)

	r.Run() // default port is :8080
}
