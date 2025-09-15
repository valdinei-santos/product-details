package vo

import "github.com/valdinei-santos/product-details/modules/product/domain/localerror"

type DescricaoProduto string

func NewDescricaoProduto(desc string) (DescricaoProduto, error) {
	if len(desc) < 5 || len(desc) > 100 {
		return "", localerror.ErrProductDescriptionInvalid
	}
	return DescricaoProduto(desc), nil
}

func (d DescricaoProduto) String() string {
	return string(d)
}
