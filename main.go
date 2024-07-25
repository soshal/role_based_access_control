// main.go
package main

import (
	"daily-api/database"
	"daily-api/handlers"
	"daily-api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to the database
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate the User model
	err = database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Define routes
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	router.POST("/sale", handlers.CreateSales)

	// Start the server
	router.Run(":8080")
}
