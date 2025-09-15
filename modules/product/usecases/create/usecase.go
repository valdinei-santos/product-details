package create

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
func (u *UseCase) Execute(in *dto.Request) (*dto.Response, error) {
	u.log.Debug("Entrou create.Execute")

	// Cria o objeto Product a partir do DTO de entrada
	p, err := entities.NewProduct(in.Nome, in.URLImagem, in.Descricao, in.Preco, in.Classificacao, in.Especificacao)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "entities.NewProduct")
		return nil, err
	}

	// Salva o produto no repositório
	err = u.repo.AddProduct(p)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.AddProduct")
		return nil, localerror.ErrProductInternal
	}
	u.log.Debug("Produto criado com sucesso", "product_id", p.ID)
	// Retorna o DTO de saída
	resp := &dto.Response{
		ID:            p.ID.String(),
		Nome:          p.Nome.String(),
		URLImagem:     p.URLImagem.String(),
		Descricao:     p.Descricao.String(),
		Preco:         p.Preco.Float64(),
		Classificacao: p.Classificacao.String(),
		Especificacao: p.Especificacao.String(),
		CreatedAt:     p.CreatedAt.String(),
		UpdatedAt:     p.UpdatedAt.String(),
	}
	return resp, nil
}
