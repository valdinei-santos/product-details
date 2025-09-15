package vo

import "github.com/valdinei-santos/product-details/modules/product/domain/localerror"

type Preco float64

func NewPreco(valor float64) (Preco, error) {
	if valor < 0 {
		return 0, localerror.ErrProductPriceNegative
	}
	return Preco(valor), nil
}

func (p Preco) Float64() float64 {
	return float64(p)
}
