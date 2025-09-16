package getall

import (
	"math"

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

// @Summary        Lista todos os produtos
// @Description    Retorna uma lista de produtos, paginada
// @Tags           produtos
// @Produce        json
// @Param          page query int false "Numero da página a ser retornada"
// @Param          size query int false "Quantidade de itens na página a ser retornada"
// @Success        200 {array} dto.ResponseManyPaginated
// @Failure        500 {string} string "Erro interno do servidor"
// @Router         /products [get]
// Execute - Executa a lógica para buscar todos os produtos
func (u *UseCase) Execute(page int, size int) (*dto.ResponseManyPaginated, error) {
	u.log.Debug("Entrou create.Execute")

	// Calcula o offset para o repositório
	offset := (page - 1) * size

	// Busca o subconjunto de produtos e o total de itens
	paginatedProducts, totalItems, err := u.repo.GetAllProducts(offset, size)
	if err != nil {
		u.log.Error("Erro ao buscar produtos: ", err)
		return nil, err
	}

	// Converte as entidades para DTOs
	productsList := make([]dto.Response, len(paginatedProducts))
	for i, p := range paginatedProducts {
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

	// Calcula o total de páginas
	totalPages := int(math.Ceil(float64(totalItems) / float64(size)))
	if totalPages == 0 && totalItems > 0 { // Lida com o caso de 1 página.
		totalPages = 1
	}

	// Constrói a resposta paginada
	result := &dto.ResponseManyPaginated{
		Products:     productsList,
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		CurrentPage:  page,
		ItemsPerPage: size,
	}

	return result, nil
}
