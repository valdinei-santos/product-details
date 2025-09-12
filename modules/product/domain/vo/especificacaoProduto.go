package vo

import "errors"

type EspecificacaoProduto string

func NewEspecificacaoProduto(desc string) (EspecificacaoProduto, error) {
	if len(desc) < 10 || len(desc) > 200 {
		return "", errors.New("a especificacao deve ter entre 10 e 200 caracteres")
	}
	return EspecificacaoProduto(desc), nil
}
