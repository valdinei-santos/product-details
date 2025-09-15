package update

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

// UseCase - Estrutura para o caso de uso de criação de produto
type UseCase struct {
	repo repository.IProductRepository // Interface do repositório para Produto
	log  logger.ILogger
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IProductRepository, l logger.ILogger) *UseCase {
	return &UseCase{
		repo: r,
		log:  l,
	}
}

// Execute - Executa a lógica de criação de um produto
func (u *UseCase) Execute(id string, in *dto.Request) (*dto.Response, error) {
	u.log.Debug("Entrou create.Execute")

	// Pega o produto no repositório pelo ID
	p, err := u.repo.GetProductByID(id)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.GetProductByID")
		return nil, localerror.ErrProductNotFound
	}

	// Altera o objeto Product a partir do Product enviado no DTO de entrada
	pNew, err := entities.UpdateProduct(id, in.Nome, in.URLImagem, in.Descricao, in.Preco, in.Classificacao, in.Especificacao, p.CreatedAt)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "entities.UpdateProduct")
		return nil, err
	}

	// Altera o produto no repositório
	err = u.repo.UpdateProduct(id, pNew)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.UpdateProduct")
		return nil, err
	}

	// Retorna o DTO de saída
	resp := &dto.Response{
		ID:            pNew.ID.String(),
		Nome:          pNew.Nome.String(),
		URLImagem:     pNew.URLImagem.String(),
		Descricao:     pNew.Descricao.String(),
		Preco:         pNew.Preco.Float64(),
		Classificacao: pNew.Classificacao.String(),
		Especificacao: pNew.Especificacao.String(),
		CreatedAt:     pNew.CreatedAt.String(),
		UpdatedAt:     pNew.UpdatedAt.String(),
	}
	return resp, nil
}
