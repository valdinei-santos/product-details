# Variables
APP_NAME=product-details

.PHONY: default run build test cover docs clean

default: build
run:
	@go run cmd/api/main.go
build:
	@go build -o $(APP_NAME) cmd/api/main.go
	@echo "Build completo: Executável da API $(APP_NAME) gerado"
test:
	go test ./...
cover:
	go test -v -cover ./...
docs:
	@swag init -g ./cmd/api/main.go -o ./docs
	@echo "Documentação gerada em internal/docs"
clean:
	@rm -f $(APP_NAME)
	@echo "Limpeza completa: Executável removido"
