package delete

import (
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

// UseCase - Estrutura para o caso de uso de delete do produto
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

// @Summary      Deleta um produto pelo ID
// @Description  Deleta um produto específico com base no ID fornecido
// @Tags         produtos
// @Produce      json
// @Param        id path string true "ID do produto a ser deletado"
// @Success      200 {object} dto.OutputDefault "Produto deletado com sucesso"}
// @Failure      404 {string} string "produto não encontrado"
// @Router       /{id} [delete]
// Execute - Executa a lógica para deletar um produto
func (u *UseCase) Execute(id string) error {
	u.log.Debug("Entrou delete.Execute")

	err := u.repo.DeleteProduct(id)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "u.repo.Delete")
		return err
	}

	return nil
}
