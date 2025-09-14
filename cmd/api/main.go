package main

import (
	"fmt"
	"io"
	"os"

	"github.com/valdinei-santos/product-details/cmd/api/routes"
	"github.com/valdinei-santos/product-details/infra/config"
	"github.com/valdinei-santos/product-details/infra/database/datafake"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Iniciando...")
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Erro ao carregar variáveis do .env: %v", err)
		os.Exit(1)
	}

	log := logger.NewSlogILogger()
	fmt.Println("Iniciou Log...")

	repoProducts, err := repository.NewProductRepo("infra/database/products.json")
	if err != nil {
		log.Debug(err.Error())
		fmt.Printf("Erro ao iniciar database: %v", err)
		os.Exit(1)
	}
	fmt.Println("Iniciou Database...")

	err = datafake.GerarProdutosFake(repoProducts, 5)
	if err != nil {
		log.Debug(err.Error())
		fmt.Printf("Erro ao gerar dados fake: %v", err)
		os.Exit(1)
	}

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
