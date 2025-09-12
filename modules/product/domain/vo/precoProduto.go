package vo

import "errors"

type Preco float64

func NewPreco(valor float64) (Preco, error) {
	if valor < 0 {
		return 0, errors.New("o valor do preço não pode ser negativo")
	}
	return Preco(valor), nil
}
