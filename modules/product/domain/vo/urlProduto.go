package vo

import "errors"

type urlProduto string

func NewUrlProduto(desc string) (urlProduto, error) {
	if len(desc) < 10 || len(desc) > 300 {
		return "", errors.New("a url deve ter entre 10 e 300 caracteres")
	}
	return urlProduto(desc), nil
}
