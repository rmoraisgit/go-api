package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080") // Run the server on port 8080
	fmt.Printf("Server is running on http://localhost:8080")
}
