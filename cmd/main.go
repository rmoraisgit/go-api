package main

import (
	"fmt"
	"go-api/internal/controllers"
	"go-api/internal/db"
	"go-api/internal/repositories"
	"go-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	_produtosRepository := repositories.NewProdutosRepository(dbConnection)
	_produtoUseCase := usecases.NewProdutosUseCase(_produtosRepository)

	produtosController :=
		controllers.NewProdutosController(_produtoUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/produtos", produtosController.ObterProdutos)
	server.GET("/produtos/:id", produtosController.ObterProdutoPorId)
	server.POST("/produtos", produtosController.CriarProduto)

	server.Run(":8080") // Run the server on port 8080
	fmt.Printf("Server is running on http://localhost:8080")
}
