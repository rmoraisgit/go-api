package usecases

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type ProdutosUseCase struct {
	_produtoRepository repositories.ProdutosRepository
	// Add any necessary fields or dependencies here
}

func NewProdutosUseCase(produtoRepository repositories.ProdutosRepository) ProdutosUseCase {
	return ProdutosUseCase{
		_produtoRepository: produtoRepository,
	}
}

func (p *ProdutosUseCase) ObterProdutos() ([]models.Product, error) {
	return p._produtoRepository.ObterProdutos()
}
