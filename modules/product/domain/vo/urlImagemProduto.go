package vo

import "github.com/valdinei-santos/product-details/modules/product/domain/localerror"

type UrlImagemProduto string

func NewUrlImagemProduto(desc string) (UrlImagemProduto, error) {
	if len(desc) < 10 || len(desc) > 300 {
		return "", localerror.ErrProductUrlImageInvalid
	}
	return UrlImagemProduto(desc), nil
}

func (u UrlImagemProduto) String() string {
	return string(u)
}
