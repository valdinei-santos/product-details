package entities

import "github.com/go-playground/validator/v10"

// Product representa a estrutura de um produto
type Product struct {
	ID            int     `json:"id"`
	Nome          string  `json:"nome"`
	URL           string  `json:"url"`
	Descricao     string  `json:"descricao"`
	Preco         float64 `json:"preco"`
	Classificacao string  `json:"classificacao"`
	Especificacao string  `json:"especificacao"`
}

// NewProduct cria uma nova instância de Product
func NewProduct(id int, nome, url, descricao string, preco float64, classificacao, especificacao string) (*Product, error) {
	p := &Product{
		ID:            id,
		Nome:          nome,
		URL:           url,
		Descricao:     descricao,
		Preco:         preco,
		Classificacao: classificacao,
		Especificacao: especificacao,
	}
	err := p.validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate - Valida os campos do Produto
func (p *Product) validate() error {
	return validator.New().Struct(p)
}
