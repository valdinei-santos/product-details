package create

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
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
func (u *UseCase) Execute(in *dto.Request) (*dto.OutputDefault, error) {
	u.log.Debug("Entrou create.Execute")

	// Cria o objeto Product a partir do DTO de entrada
	p := &entities.Product{
		ID:            in.ID,
		Nome:          in.Nome,
		URL:           in.URL,
		Descricao:     in.Descricao,
		Preco:         in.Preco,
		Classificacao: in.Classificacao,
		Especificacao: in.Especificacao,
	}

	// Salva o produto no repositório
	err := u.repo.AddProduct(p)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.AddProduct")
		return nil, err
	}

	// Retorna a resposta padrão
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Produto inserido com sucesso",
	}
	return result, nil
}
