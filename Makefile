# Variables
APP_NAME=product-details

.PHONY: help run build test cover docs clean

default: help
help: ## Exibe esta mensagem de ajuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
run: ## Roda o projeto
	@go run cmd/api/main.go
build: ## Cria o executável da aplicação
	@go build -o $(APP_NAME) cmd/api/main.go
	@echo "Build completo: Executável da API $(APP_NAME) gerado"
test: ## Executa os test automatizados da aplicação
	go test ./...
cover: ## Mostra a cobertura de testes da aplicação
	go test -v -cover ./...
docs: ## Gera a documentação OpenAPI (Swagger) dos endpoints da aplicação
	@swag init -g ./cmd/api/main.go -o ./docs
	@echo "Documentação gerada em internal/docs"
clean: ## Remove o executável gerado
	@rm -f $(APP_NAME)
	@echo "Limpeza completa: Executável removido"
