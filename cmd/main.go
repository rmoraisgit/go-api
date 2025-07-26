package main

import (
	"fmt"
	"go-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	produtosController := controllers.NewProdutosController()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/produtos", produtosController.ObterProdutos)

	server.Run(":8080") // Run the server on port 8080
	fmt.Printf("Server is running on http://localhost:8080")
}
