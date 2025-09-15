package vo

import "github.com/valdinei-santos/product-details/modules/product/domain/localerror"

type EspecificacaoProduto string

func NewEspecificacaoProduto(desc string) (EspecificacaoProduto, error) {
	if len(desc) < 10 || len(desc) > 200 {
		return "", localerror.ErrProductSpecificationInvalid
	}
	return EspecificacaoProduto(desc), nil
}

func (e EspecificacaoProduto) String() string {
	return string(e)
}
