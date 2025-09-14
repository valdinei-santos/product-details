package routes

import (
	_ "github.com/valdinei-santos/product-details/cmd/api/docs" // swagger docs

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

func InitRoutes(router *gin.RouterGroup, log logger.ILogger, repoProducts repository.IProductRepository) {

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	v1 := router.Group("/api/v1")
	prod := v1.Group("/products")

	prod.GET("/", func(c *gin.Context) {
		log.Info("### Start endpoint GET /api/v1/products")
		product.StartGetAll(log, c, repoProducts)
	})

	prod.POST("/", func(c *gin.Context) {
		log.Info("### Start endpoint POST /api/v1/products")
		product.StartCreate(log, c, repoProducts)
	})

	prod.GET("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint GET /api/v1/products/:id")
		product.StartGet(log, c, repoProducts)
	})

	// @Summary      Atualiza um produto
	// @Description  Atualiza um produto existente com base no ID
	// @Tags         produtos
	// @Accept       json
	// @Produce      json
	// @Param        id   path      string  true  "Product ID"
	// @Param        produto  body      dto.Request  true  "Dados do produto para atualização"
	// @Success      200 {object}  dto.OutputDefault
	// @Failure      400 {object}  dto.OutputDefault
	// @Router       /products/{id} [put]
	prod.PUT("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint PUT /api/v1/products/:id")
		product.StartUpdate(log, c, repoProducts)
	})

	prod.DELETE("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint DELETE /api/v1/products/:id")
		product.StartDelete(log, c, repoProducts)
	})

	prod.GET("/compare", func(c *gin.Context) {
		log.Info("### Start endpoint GET /api/v1/products/compare")
		product.StartCompare(log, c, repoProducts)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

}
