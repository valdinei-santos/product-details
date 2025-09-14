package vo

import "errors"

type DescricaoProduto string

func NewDescricaoProduto(desc string) (DescricaoProduto, error) {
	if len(desc) < 5 || len(desc) > 100 {
		return "", errors.New("a descricao deve ter entre 5 e 100 caracteres")
	}
	return DescricaoProduto(desc), nil
}

func (d DescricaoProduto) String() string {
	return string(d)
}
