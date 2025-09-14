package vo

import "errors"

type ClassificacaoProduto string

func NewClassificacaoProduto(desc string) (ClassificacaoProduto, error) {
	if len(desc) < 3 || len(desc) > 50 {
		return "", errors.New("a classificacao deve ter entre 3 e 50 caracteres")
	}
	return ClassificacaoProduto(desc), nil
}

func (c ClassificacaoProduto) String() string {
	return string(c)
}
