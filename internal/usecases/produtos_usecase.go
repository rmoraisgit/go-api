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

func (p *ProdutosUseCase) CriarProduto(produto models.Product) (models.Product, error) {
	return p._produtoRepository.CriarProduto(produto)
}

func (p *ProdutosUseCase) ObterProdutoPorId(id int) (models.Product, error) {
	return p._produtoRepository.ObterProdutoPorId(id)
}
