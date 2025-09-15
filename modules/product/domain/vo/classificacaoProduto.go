package vo

import "github.com/valdinei-santos/product-details/modules/product/domain/localerror"

type ClassificacaoProduto string

func NewClassificacaoProduto(desc string) (ClassificacaoProduto, error) {
	if len(desc) < 3 || len(desc) > 50 {
		return "", localerror.ErrProductClassificationInvalid
	}
	return ClassificacaoProduto(desc), nil
}

func (c ClassificacaoProduto) String() string {
	return string(c)
}
