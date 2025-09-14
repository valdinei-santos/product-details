package vo

import "errors"

type UrlProduto string

func NewUrlProduto(desc string) (UrlProduto, error) {
	if len(desc) < 10 || len(desc) > 300 {
		return "", errors.New("a url deve ter entre 10 e 300 caracteres")
	}
	return UrlProduto(desc), nil
}

func (u UrlProduto) String() string {
	return string(u)
}
