package vo

import "errors"

type NomeProduto string

func NewNomeProduto(nome string) (NomeProduto, error) {
	if len(nome) < 3 || len(nome) > 50 {
		return "", errors.New("o nome deve ter entre 3 e 50 caracteres")
	}
	return NomeProduto(nome), nil
}

func (n NomeProduto) String() string {
	return string(n)
}
