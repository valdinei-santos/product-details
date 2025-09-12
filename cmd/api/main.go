package main

import (
	"fmt"
	"io"
	"log"

	"github.com/valdinei-santos/product-details/cmd/api/routes"
	"github.com/valdinei-santos/product-details/infra/config"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Iniciando...")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis do .env: %v", err)
	}

	log := logger.NewSlogLogger()
	fmt.Println("Iniciou Log...")

	repoProducts, err := repository.NewProductRepo("infra/database/products.json")
	if err != nil {
		log.Debug(err.Error())
		return
	}
	fmt.Println("Iniciou Database...")

	gin.DefaultWriter = io.Discard // Desabilita o log padrão do gin jogando para o io.Discard
	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.InitRoutes(&router.RouterGroup, log, repoProducts)

	log.Info("start product-details", "PORT:", config.Port)
	err = router.Run(":" + config.Port)
	if err != nil {
		fmt.Printf("Erro ao iniciar a API na porta %v: %v", config.Port, err)
		log.Error("Erro ao iniciar a API na porta " + config.Port + " - " + err.Error())
	}
	router.Run(":" + config.Port)
}
