package main

import (
	controllers "go-clean-app/controller"
	mongoConnection "go-clean-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoConnection.ConnectMongo()
	r := gin.Default()

	// Route to get dummy user
	r.GET("/user", controllers.CreateAndStoreDummyUser)
	r.GET("/users", controllers.GetAllUsers)

	r.Run() // default port is :8080
}
