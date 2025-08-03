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

func (p *ProdutosRepository) ObterProdutos() ([]models.Product, error) {

	query := "SELECT id, nome, preco FROM produtos"

	rows, err := p._connectionDb.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Nome,
			&productObj.Preco)

		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (p *ProdutosRepository) CriarProduto(produto models.Product) (models.Product, error) {

	query := "INSERT INTO produtos (nome, preco) VALUES ($1, $2) RETURNING id, nome, preco"

	err := p._connectionDb.QueryRow(query, produto.Nome, produto.Preco).Scan(&produto.Id, &produto.Nome, &produto.Preco)

	if err != nil {
		fmt.Println("Error inserting product:", err)
		return models.Product{}, err
	}

	return produto, nil
}

func (p *ProdutosRepository) ObterProdutoPorId(id string) (models.Product, error) {

	query := "SELECT id, nome, preco FROM produtos WHERE id = $1"

	var produto models.Product

	err := p._connectionDb.QueryRow(query, id).Scan(&produto.Id, &produto.Nome, &produto.Preco)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, nil // Produto não encontrado
		}
		fmt.Println("Error fetching product by ID:", err)
		return models.Product{}, nil
	}

	return produto, nil
}
