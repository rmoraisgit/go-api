package controllers

import (
	"go-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type produtosController struct {
	// Add any necessary fields or dependencies here
}

func NewProdutosController() produtosController {
	return produtosController{}
}

func (p *produtosController) ObterProdutos(ctx *gin.Context) {
	// Implement the logic to retrieve products
	// This is a placeholder for demonstration purposes
	products := []models.Product{
		{Id: 1, Nome: "Produto A", Preco: 10.0},
		{Id: 2, Nome: "Produto B", Preco: 20.0},
	}

	ctx.JSON(http.StatusOK, products)
}
