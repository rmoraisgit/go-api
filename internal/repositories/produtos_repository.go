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
