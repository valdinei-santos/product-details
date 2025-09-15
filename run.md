# API REST product-details

Esta é uma API REST feita em Golang para gerenciar produtos, permitindo a comparação de múltiplos produtos.

## Pré-requisitos

- Sistema operacional Linux.
- Ferramenta `make` do Linux instalada na sua distribuição.
- Go versão 1.22 ou superior instalado.

## 1. Instalação

Clone o repositório e navegue até a pasta do projeto:
```bash
git clone https://github.com/valdinei-santos/product-details.git
cd product-details
```

Para baixar todas as dependências do projeto, use o comando:
```bash
go mod tidy
```

## 2. Configuração
Para rodar a API, você precisa criar um arquivo de variáveis de ambiente .env. Copie o arquivo de exemplo env.example, nesse momento você só precisa da definição da PORTA que a API vai rodar:
```bash
cp .env.example .env
```

## 3. Execução

Para compilar a API, use o seguinte comando.
```bash
make build
```

Para rodar a API, você simplementes executa o arquivo gerado no passo anterior:
```bash
./product-details
```

## 4. Testes
Para rodar os testes unitários e de integração do projeto, siga os passos abaixo:

1. Navegue até o diretório do projeto. Caso a API esteja rodando você precisa parar ela com CTRL+C:
```bash
cd product-details
```

2. Execute todos os testes:
```bash
make test
```

3. Para rodar um arquivo de teste específico, use o comando **go test "nome-do-arquivo**, conforme abaixo:
```bash
go test modules/product/usecases/delete/usecase_test.go
```

4. Para rodar um caso de teste específico, use o comando **go test -run "nome-do-caso-de-teste"**, conforme abaixo:
```bash
go test -run "Deve retornar sucesso ao excluir um produto" modules/product/usecases/delete/usecase_test.go
```

5. Para ver a cobertura dos testes na aplicação:
```bash
make cover
``` 

### Estrutura de Testes
O projeto inclui testes automatizados para os seguintes pacotes:

- **cmd/api/routes**: Faz os testes de integração de todos os endpoints.
- **modules/product/usecases/compare**: Faz testes de unidade do usecase **compare** 
- **modules/product/usecases/create**: Faz testes de unidade do usecase **create**
- **modules/product/usecases/delete**: Faz testes de unidade do usecase **delete**
- **modules/product/usecases/get**: Faz testes de unidade do usecase **get**
- **modules/product/usecases/getall**: Faz testes de unidade do usecase **getall**
- **modules/product/usecases/update**: Faz testes de unidade do usecase **update**
