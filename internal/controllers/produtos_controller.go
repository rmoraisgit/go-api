package controllers

import (
	"go-api/internal/models"
	"go-api/internal/usecases"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ProdutosController struct {
	_produtoUseCase usecases.ProdutosUseCase
}

func NewProdutosController(produtoUseCase usecases.ProdutosUseCase) ProdutosController {
	return ProdutosController{
		_produtoUseCase: produtoUseCase,
	}
}

func (p *ProdutosController) ObterProdutos(ctx *gin.Context) {

	products, err := p._produtoUseCase.ObterProdutos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProdutosController) CriarProduto(ctx *gin.Context) {

	var produto models.Product

	err := ctx.ShouldBindJSON(&produto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	produtoCriado, err := p._produtoUseCase.CriarProduto(produto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, produtoCriado)
}
