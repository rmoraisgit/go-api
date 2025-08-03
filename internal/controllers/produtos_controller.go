package controllers

import (
	"go-api/internal/models"
	"go-api/internal/usecases"
	"strconv"

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

	var produto models.Produto

	err := ctx.ShouldBindJSON(&produto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	produtoCriado, err := p._produtoUseCase.CriarProduto(produto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, produtoCriado)
}

func (p *ProdutosController) ObterProdutoPorId(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo id é requerido"})
		return
	}

	idRecebidoRequisicao, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo id deve ser um número"})
		return
	}

	produto, err := p._produtoUseCase.ObterProdutoPorId(idRecebidoRequisicao)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	produtoVazio := models.Produto{}

	if produto == produtoVazio {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, produto)
}
