package get

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

// UseCase - Struct do  caso de uso
type UseCase struct {
	repo repository.IProductRepository // Interface do repositório para Produto
	log  logger.ILogger                // Interface do log
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IProductRepository, l logger.ILogger) *UseCase {
	return &UseCase{
		repo: r,
		log:  l,
	}
}

// Execute - Executa a lógica de busca de um produto
func (u *UseCase) Execute(id string) (*dto.Response, error) {
	u.log.Debug("Entrou get.Execute")

	// Pega o produto no repositório pelo ID
	p, err := u.repo.GetProductByID(id)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.GetProductByID")
		return nil, err
	}

	// Transforma a entidade Product no DTO Response
	result := &dto.Response{
		ID:            p.ID.String(),
		Nome:          p.Nome.String(),
		URLImagem:     p.URLImagem.String(),
		Descricao:     p.Descricao.String(),
		Preco:         p.Preco.Float64(),
		Classificacao: p.Classificacao.String(),
		Especificacao: p.Especificacao.String(),
	}
	return result, nil
}
