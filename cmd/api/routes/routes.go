package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

func InitRoutes(router *gin.RouterGroup, log logger.Logger, repoProducts repository.IProductRepository) {

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
	api := router.Group("/api")
	prod := api.Group("/products")

	prod.GET("/", func(c *gin.Context) {
		log.Info("### Start endpoint GET /api/products")
		product.StartGetAll(log, c, repoProducts)
	})

	prod.POST("/", func(c *gin.Context) {
		log.Info("### Start endpoint POST /api/products")
		product.StartCreate(log, c, repoProducts)
	})

	prod.GET("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint GET /api/products/:id")
		product.StartGet(log, c, repoProducts)
	})

	prod.PUT("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint PUT /api/products/:id")
		product.StartUpdate(log, c, repoProducts)
	})

	prod.DELETE("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint DELETE /api/products/:id")
		product.StartDelete(log, c, repoProducts)
	})
}
