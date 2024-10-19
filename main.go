package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// // Configure CORS middleware
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your allowed origin(s)
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	// 	AllowCredentials: true,
	// }))

	router.Use(Middleware("*", "ecom-system"))

	// Define a GET endpoint
	router.GET("/get-endpoint", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET request successful",
		})
	})

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

const (
	MerchantPortalExcludeURL = "http://localhost:3000"
)

// Middleware handles cross origin request access
func Middleware(origins, service string) gin.HandlerFunc {
	if strings.TrimSpace(origins) == "" {
		logrus.Fatal("CORS origin must be provided for server")
	}

	return func(c *gin.Context) {
		reqOrigin := c.Request.Header.Get("Origin")
		if reqOrigin == MerchantPortalExcludeURL {
			c.Writer.Header().Set("Access-Control-Allow-Origin", MerchantPortalExcludeURL)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origins)
		}
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
