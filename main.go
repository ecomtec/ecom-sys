package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define your API routes
	router.GET("/hello", helloHandler)
	router.POST("/users", createUserHandler)
	router.GET("/users/:id", getUserHandler)
	router.PUT("/users/:id", updateUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)

	// Start the server
	router.Run(":80")
	router.Routes()
}

func helloHandler(c *gin.Context) {
	// Return a simple message
	c.JSON(200, gin.H{
		"message": "Hello, Suhail.....!",
	})
}

func createUserHandler(c *gin.Context) {
	// Create a new user
	// ...
}

func getUserHandler(c *gin.Context) {
	// Get a user by ID
	// ...
}

func updateUserHandler(c *gin.Context) {
	// Update a user by ID
	// ...
}

func deleteUserHandler(c *gin.Context) {
	// Delete a user by ID
	// ...
}
