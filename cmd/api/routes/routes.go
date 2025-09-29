// Pacote routes configura as rotas da API.
package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/valdinei-santos/product-details/docs" // swagger docs
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

func InitRoutes(router *gin.RouterGroup, log logger.ILogger, repoProducts repository.IProductRepository) {

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8888", "http://127.0.0.1:8888"}, // Para liberar o Swagger
		// OU use AllowAllOrigins: true para permitir TUDO (apenas para DEV/TESTE)

		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Limita tempo navegador pode armazenar em cache as informações do preflight
	}))
	// ---------------------------

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/api/v1")
	prod := v1.Group("/products")

	prod.GET("/", func(c *gin.Context) {
		log.Info("### Start endpoint " + c.Request.Method + " " + c.Request.URL.Path)
		product.StartGetAll(log, c, repoProducts)
	})

	prod.POST("/", func(c *gin.Context) {
		log.Info("### Start endpoint " + c.Request.Method + " " + c.Request.URL.Path)
		product.StartCreate(log, c, repoProducts)
	})

	prod.GET("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint " + c.Request.Method + " " + c.Request.URL.Path)
		product.StartGet(log, c, repoProducts)
	})

	prod.PUT("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint " + c.Request.Method + " " + c.Request.URL.Path)
		product.StartUpdate(log, c, repoProducts)
	})

	prod.DELETE("/:id", func(c *gin.Context) {
		log.Info("### Start endpoint " + c.Request.Method + " " + c.Request.URL.Path)
		product.StartDelete(log, c, repoProducts)
	})

	prod.GET("/compare", func(c *gin.Context) {
		log.Info("### Start endpoint " + c.Request.Method + " " + c.Request.URL.Path)
		product.StartCompare(log, c, repoProducts)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

}
