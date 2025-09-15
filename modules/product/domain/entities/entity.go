package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/domain/vo"
)

// Product representa a estrutura de um produto
type Product struct {
	ID            uuid.UUID
	Nome          vo.NomeProduto
	URLImagem     vo.UrlImagemProduto
	Descricao     vo.DescricaoProduto
	Preco         vo.Preco
	Classificacao vo.ClassificacaoProduto
	Especificacao vo.EspecificacaoProduto
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// NewProduct - cria uma nova instância de Product
func NewProduct(nome, urlImagem, descricao string, preco float64, classificacao, especificacao string) (*Product, error) {
	uuidVO, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	nomeVO, err := vo.NewNomeProduto(nome)
	if err != nil {
		return nil, err
	}
	urlImagemVO, err := vo.NewUrlImagemProduto(urlImagem)
	if err != nil {
		return nil, err
	}
	descricaoVO, err := vo.NewDescricaoProduto(descricao)
	if err != nil {
		return nil, err
	}
	precoVO, err := vo.NewPreco(preco)
	if err != nil {
		return nil, err
	}
	classificacaoVO, err := vo.NewClassificacaoProduto(classificacao)
	if err != nil {
		return nil, err
	}
	especificacaoVO, err := vo.NewEspecificacaoProduto(especificacao)
	if err != nil {
		return nil, err
	}

	p := &Product{
		ID:            uuidVO,
		Nome:          nomeVO,
		URLImagem:     urlImagemVO,
		Descricao:     descricaoVO,
		Preco:         precoVO,
		Classificacao: classificacaoVO,
		Especificacao: especificacaoVO,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = p.validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// UpdateProduct - altera uma instância de Product
func UpdateProduct(id, nome, urlImagem, descricao string, preco float64, classificacao, especificacao string, createdAt time.Time) (*Product, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, localerror.ErrProductIDInvalid
	}
	nomeVO, err := vo.NewNomeProduto(nome)
	if err != nil {
		return nil, err
	}
	urlImagemVO, err := vo.NewUrlImagemProduto(urlImagem)
	if err != nil {
		return nil, err
	}
	descricaoVO, err := vo.NewDescricaoProduto(descricao)
	if err != nil {
		return nil, err
	}
	precoVO, err := vo.NewPreco(preco)
	if err != nil {
		return nil, err
	}
	classificacaoVO, err := vo.NewClassificacaoProduto(classificacao)
	if err != nil {
		return nil, err
	}
	especificacaoVO, err := vo.NewEspecificacaoProduto(especificacao)
	if err != nil {
		return nil, err
	}

	p := &Product{
		ID:            idUUID,
		Nome:          nomeVO,
		URLImagem:     urlImagemVO,
		Descricao:     descricaoVO,
		Preco:         precoVO,
		Classificacao: classificacaoVO,
		Especificacao: especificacaoVO,
		CreatedAt:     createdAt,
		UpdatedAt:     time.Now(),
	}
	err = p.validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate - Valida os campos do Produto
func (p *Product) validate() error {
	return validator.New().Struct(p)
}
