package get

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

// UseCase - Estrutura para o caso de uso de criação de produto
type UseCase struct {
	repo repository.IProductRepository // Interface do repositório para Produto
	log  logger.Logger
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IProductRepository, l logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  l,
	}
}

// Execute - Executa a lógica de criação de um produto
func (u *UseCase) Execute(id int) (*dto.Response, error) {
	u.log.Debug("Entrou get.Execute")

	// Salva o produto no repositório
	p, err := u.repo.GetProductByID(id)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.GetProductByID")
		return nil, err
	}

	result := &dto.Response{
		ID:            p.ID,
		Nome:          p.Nome,
		URL:           p.URL,
		Descricao:     p.Descricao,
		Preco:         p.Preco,
		Classificacao: p.Classificacao,
		Especificacao: p.Especificacao,
	}
	return result, nil
}
