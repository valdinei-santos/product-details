package vo

import "github.com/valdinei-santos/product-details/modules/product/domain/localerror"

type NomeProduto string

func NewNomeProduto(nome string) (NomeProduto, error) {
	if len(nome) < 3 || len(nome) > 50 {
		return "", localerror.ErrProductNameInvalid
	}
	return NomeProduto(nome), nil
}

func (n NomeProduto) String() string {
	return string(n)
}
