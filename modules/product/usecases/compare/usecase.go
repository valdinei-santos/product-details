package compare

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/infra/dto"
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

// @Summary      Compara múltiplos produtos por uma lista de IDs
// @Description  Retorna uma lista de produtos com base nos IDs fornecidos em um formato de lista separada por vírgulas
// @Tags         produtos
// @Produce      json
// @Param        ids query string true "IDs dos produtos a serem comparados (ex: 0d605862-91e8-11f0-9140-00155d6d572f,034aeff8-90ea-11f0-95f2-00155d6d5ec0,034afa35-90ea-11f0-95f2-00155d6d5ec0,034b11f9-90ea-11f0-95f2-00155d6d5ec0)"
// @Success      200 {array} dto.ResponseManyPaginated
// @Failure      400 {object} string "Requisição inválida"
// @Router       /compare [get]
// Execute - Executa a lógica de comparação de produtos
func (u *UseCase) Execute(ids []string) (*dto.ResponseMany, error) {
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
			URLImagem:     p.URLImagem.String(),
			Descricao:     p.Descricao.String(),
			Preco:         p.Preco.Float64(),
			Classificacao: p.Classificacao.String(),
			Especificacao: p.Especificacao.String(),
		}
	}

	result := &dto.ResponseMany{
		Products: productsList,
	}

	return result, nil
}
