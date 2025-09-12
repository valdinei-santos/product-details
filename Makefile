# Variables
APP_NAME=product-details

# Tasks
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
