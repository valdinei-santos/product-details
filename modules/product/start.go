package product

import (
	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/infra/controller"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/create"
	"github.com/valdinei-santos/product-details/modules/product/usecases/delete"
	"github.com/valdinei-santos/product-details/modules/product/usecases/get"
	getall "github.com/valdinei-santos/product-details/modules/product/usecases/get-all"
	"github.com/valdinei-santos/product-details/modules/product/usecases/update"
)

// StartCreate - Metodo onde instanciamentos as dependencias e chamamos o controller
func StartCreate(log logger.Logger, ctx *gin.Context, repo repository.IProductRepository) {
	log.Debug("Entrou product.StartCreate")
	u := create.NewUseCase(repo, log)
	controller.Create(log, ctx, u)
}

// StartDelete - Metodo onde instanciamentos as dependencias e chamamos o controller
func StartDelete(log logger.Logger, ctx *gin.Context, repo repository.IProductRepository) {
	log.Debug("Entrou product.StartDelete")
	u := delete.NewUseCase(repo, log)
	controller.Delete(log, ctx, u)
}

// StartGet - Metodo onde instanciamentos as dependencias e chamamos o controller
func StartGet(log logger.Logger, ctx *gin.Context, repo repository.IProductRepository) {
	log.Debug("Entrou product.StartGet")
	u := get.NewUseCase(repo, log)
	controller.Get(log, ctx, u)
}

// StartGetAll - Metodo onde instanciamentos as dependencias e chamamos o controller
func StartGetAll(log logger.Logger, ctx *gin.Context, repo repository.IProductRepository) {
	log.Debug("Entrou usecases.StartGetAll")
	u := getall.NewUseCase(repo, log)
	controller.GetAll(log, ctx, u)
}

// StartUpdate - Metodo onde instanciamentos as dependencias e chamamos o controller
func StartUpdate(log logger.Logger, ctx *gin.Context, repo repository.IProductRepository) {
	log.Debug("Entrou product.StartUpdate")
	u := update.NewUseCase(repo, log)
	controller.Update(log, ctx, u)
}
