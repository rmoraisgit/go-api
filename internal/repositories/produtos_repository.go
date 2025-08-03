package repositories

import (
	"database/sql"
	"fmt"
	"go-api/internal/models"
)

type ProdutosRepository struct {
	_connectionDb *sql.DB
}

func NewProdutosRepository(connectionDb *sql.DB) ProdutosRepository {
	return ProdutosRepository{
		_connectionDb: connectionDb,
	}
}

func (p *ProdutosRepository) ObterProdutos() ([]models.Produto, error) {

	query := "SELECT id, nome, preco FROM produtos"

	rows, err := p._connectionDb.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return []models.Produto{}, err
	}

	var ProdutoList []models.Produto
	var ProdutoObj models.Produto

	for rows.Next() {
		err = rows.Scan(
			&ProdutoObj.Id,
			&ProdutoObj.Nome,
			&ProdutoObj.Preco)

		if err != nil {
			fmt.Println(err)
			return []models.Produto{}, err
		}

		ProdutoList = append(ProdutoList, ProdutoObj)
	}

	rows.Close()

	return ProdutoList, nil
}

func (p *ProdutosRepository) CriarProduto(produto models.Produto) (models.Produto, error) {

	query := "INSERT INTO produtos (nome, preco) VALUES ($1, $2) RETURNING id, nome, preco"

	err := p._connectionDb.QueryRow(query, produto.Nome, produto.Preco).Scan(&produto.Id, &produto.Nome, &produto.Preco)

	if err != nil {
		fmt.Println("Error inserting Produto:", err)
		return models.Produto{}, err
	}

	return produto, nil
}

func (p *ProdutosRepository) ObterProdutoPorId(id int) (models.Produto, error) {

	query := "SELECT id, nome, preco FROM produtos WHERE id = $1"

	var produto models.Produto

	err := p._connectionDb.QueryRow(query, id).Scan(&produto.Id, &produto.Nome, &produto.Preco)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Produto{}, nil // Produto não encontrado
		}
		fmt.Println("Error fetching Produto by ID:", err)
		return models.Produto{}, nil
	}

	return produto, nil
}
