package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define your API routes
	router.GET("/hello", helloHandler)
	router.POST("/users", createUserHandler)
	router.POST("/login", Login)
	router.GET("/users/:id", getUserHandler)
	router.PUT("/users/:id", updateUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)

	// Start the server
	router.Run(":8081")
	router.Routes()
}

type LoginPayload struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func Login(c *gin.Context) {

	data := []LoginPayload{
		LoginPayload{
			UserName: "suhail",
			PassWord: "123",
		},
		LoginPayload{
			UserName: "sinan",
			PassWord: "345"},
	}

	fmt.Println(data)

	var creds LoginPayload

	// Parse JSON request body into creds struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println("Json incommming Data ", creds)

	var peopleSlice []LoginPayload

	for _, pp := range data {

		fmt.Printf(pp.PassWord)

		peopleSlice = append(peopleSlice, creds)
	}

	// Print the slice of people
	fmt.Println("People in slice:")
	for _, p := range peopleSlice {
		fmt.Printf("Name: %s ", p.UserName)
		fmt.Printf("Age:", p.PassWord)
	}

	// Validate credentials (dummy check for demonstration)
	if creds.UserName == "suhail" && creds.PassWord == "123" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

	// Create a new user
	// ...
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
