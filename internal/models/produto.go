package models

type Produto struct {
	Id    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}
