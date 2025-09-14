package compare

import (
	"github.com/valdinei-santos/product-details/infra/logger"
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

// Execute - Executa a lógica de comparação de produtos
func (u *UseCase) Execute(ids []string) (*dto.ProductsResponse, error) {
	u.log.Debug("Entrou get.Execute")

	// Pega os produtos no repositório pelos IDs
	products, err := u.repo.GetManyProductByIDs(ids)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.GetManyProductByIDs")
		return nil, err
	}

	productsList := make([]dto.Response, len(products))
	for i, p := range products {
		productsList[i] = dto.Response{
			ID:            p.ID.String(),
			Nome:          p.Nome.String(),
			URL:           p.URL.String(),
			Descricao:     p.Descricao.String(),
			Preco:         p.Preco.Float64(),
			Classificacao: p.Classificacao.String(),
			Especificacao: p.Especificacao.String(),
		}
	}

	result := &dto.ProductsResponse{
		Products: productsList,
	}

	return result, nil
}
