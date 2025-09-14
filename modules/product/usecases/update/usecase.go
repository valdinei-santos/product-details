package update

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
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
func (u *UseCase) Execute(id string, in *dto.Request) (*dto.OutputDefault, error) {
	u.log.Debug("Entrou create.Execute")

	// Cria o objeto Product a partir do DTO de entrada
	p, err := entities.NewProduct(in.Nome, in.URL, in.Descricao, in.Preco, in.Classificacao, in.Especificacao)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "entities.NewProduct")
		return nil, err
	}

	// Altera o produto no repositório
	err = u.repo.UpdateProduct(id, p)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.UpdateProduct")
		return nil, err
	}

	// Retorna a resposta padrão
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Produto alterado com sucesso",
	}
	return result, nil
}
