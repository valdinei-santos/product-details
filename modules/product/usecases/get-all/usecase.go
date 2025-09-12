package getall

import (
	"math"

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

// Execute - Executa a lógica para buscar todos os produtos
func (u *UseCase) Execute(page int, size int) (*dto.ProductsPaginatedResponse, error) {
	u.log.Debug("Entrou create.Execute")

	// Calcula o offset para o repositório
	offset := (page - 1) * size

	// Busca o subconjunto de produtos e o total de itens
	paginatedProducts, totalItems := u.repo.GetAllProducts(offset, size)

	// Converte as entidades para DTOs
	productsList := make([]dto.Response, len(paginatedProducts))
	for i, p := range paginatedProducts {
		productsList[i] = dto.Response{
			ID:            p.ID,
			Nome:          p.Nome,
			URL:           p.URL,
			Descricao:     p.Descricao,
			Preco:         p.Preco,
			Classificacao: p.Classificacao,
			Especificacao: p.Especificacao,
		}
	}

	// 3. Calcula o total de páginas
	totalPages := int(math.Ceil(float64(totalItems) / float64(size)))
	if totalPages == 0 && totalItems > 0 { // Lida com o caso de 1 página.
		totalPages = 1
	}

	// 4. Constrói a resposta paginada
	result := &dto.ProductsPaginatedResponse{
		Products:     productsList,
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		CurrentPage:  page,
		ItemsPerPage: size,
	}

	/* productsList := []dto.Response{}
	products := u.repo.GetAllProducts()

	for _, p := range products {
		result := &dto.Response{
			ID:            p.ID,
			Nome:          p.Nome,
			URL:           p.URL,
			Descricao:     p.Descricao,
			Preco:         p.Preco,
			Classificacao: p.Classificacao,
			Especificacao: p.Especificacao,
		}
		productsList = append(productsList, *result)
	}
	result := &dto.ProductsResponse{
		Products: productsList,
	} */
	return result, nil
}
