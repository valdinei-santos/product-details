# API product-details

## Descrição Técnica

A API **product-details** foi criada com o intuito de demonstrarmos uma simples implementação de API Rest usando os conceitos de Arquitetura Limpa e DDD com Golang.
Para facilitar os testes o repositório usado é um arquivo em JSON, podendo outros tipos de repositórios serem adicionados posteriormente de forma totalmente desacoplada.
Na inicialização da API 5 produtos são criados por uma função Fake, para facilitar os testes.

## Instruções para Rodar o Programa

### Pré-requisitos

- Sistema operacional Linux 
- Go instalado (versão 1.22 ou superior).

### Passos

1. Clone o repositório:
   ```bash
   git clone https://github.com/valdinei-santos/product-details.git
   cd product-details
   ```

2. Crie o arquivo **.env** com o valor da chave PORT. Você também pode renomear o arquivo .env.exemplo para .env:
   ```bash
   echo "PORT=8888" > .env 
   ```
   Confira a criação do arquivo:
   ```bash
   cat .env
   ```

3. Compile o programa:
   ```bash
   make build
   ```
   Saída:
   ```bash
   Build completo: Executável da API product-details gerado
   ```

4. Execute a API rodando o arquivo gerado no Build:
   ```bash
   ./product-details
   ```

   Saída:
   ```bash
   Iniciando...
   Iniciou Log...
   Iniciou Database...
   Gerando produtos fake caso tenha menos que 5 produtos...
   {"time":"2025-09-14T00:54:33.731833513-03:00","level":"INFO","msg":"start product-details","PORT:":"8888"}
   ```

## Instruções para Rodar os Testes

### Passos

1. Navegue até o diretório do projeto. Caso a API esteja rodando você precisa parar ela com CTRL+C.:
   ```bash
   cd product-details
   ```

2. Execute todos os testes:
   ```bash
   make test
   ```

   Exemplo de saída:
   ```bash
   go test ./...
   ?       github.com/valdinei-santos/product-details/cmd/api      [no test files]
   ok      github.com/valdinei-santos/product-details/cmd/api/routes       (cached)
   ?       github.com/valdinei-santos/product-details/infra/config [no test files]
   ?       github.com/valdinei-santos/product-details/infra/database/datafake      [no test files]
   ?       github.com/valdinei-santos/product-details/infra/logger [no test files]
   ?       github.com/valdinei-santos/product-details/modules/product      [no test files]
   ?       github.com/valdinei-santos/product-details/modules/product/domain/entities      [no test files]
   ?       github.com/valdinei-santos/product-details/modules/product/domain/vo    [no test files]
   ?       github.com/valdinei-santos/product-details/modules/product/dto  [no test files]
   ?       github.com/valdinei-santos/product-details/modules/product/infra/controller     [no test files]
   ?       github.com/valdinei-santos/product-details/modules/product/infra/repository     [no test files]
   ok      github.com/valdinei-santos/product-details/modules/product/usecases/compare     (cached)
   ok      github.com/valdinei-santos/product-details/modules/product/usecases/create      (cached)
   ok      github.com/valdinei-santos/product-details/modules/product/usecases/delete      (cached)
   ok      github.com/valdinei-santos/product-details/modules/product/usecases/get (cached)
   ok      github.com/valdinei-santos/product-details/modules/product/usecases/get-all     (cached)
   ok      github.com/valdinei-santos/product-details/modules/product/usecases/update      (cached)
   ```

3. Para rodar um arquivo de teste específico, use o comando **go test**, conforme abaixo:
   ```bash
   go test caminho/do/arquivo_test.go
   
   Exemplo:
   go test modules/product/usecases/delete/usecase_test.go
   ```

4. Para rodar um caso de teste específico, use o comando **go test -run "nome-do-caso-de-teste"**, conforme abaixo:
   ```bash
   go test -run "Deve retornar sucesso ao excluir um produto" modules/product/usecases/delete/usecase_test.g
   ```
   Exemplo:
   ```bash
   go test modules/product/usecases/delete/usecase_test.go
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


## Detalhes de Implementação

### Estrutura do Projeto

- **cmd/api/**: Contém o ponto de entrada principal da API (`main.go`). Onde o programa inicia as estruturas básicas de config, log, database, framework [GIN](https://gin-gonic.com/) e alocada a porta informada no .env para uso da API.
- **cmd/api/routes**: Pacote com a definição dos **endpoints** disponíveis pela API.
- **infra**: Pasta com os pacotes da camada de infraestrutura usados no contexto geral da API (config, database e log).
- **modules**: Pasta com todos os recursos da API, que nesse caso é apenas **product**.
- **modules/product**: Pasta principal do recurso **product**.
- **modules/product/domain**: Pasta da camada de dominio (Clean Arch) do recurso.
- **modules/product/domain/entities**: Pasta com o pacote das entidades do recurso.
- **modules/product/domain/vo**: Pasta com o pacote dos Value Objects definidos.
- **modules/product/dto**: Pasta com o pacote de DTOs definidos para o recurso.
- **modules/product/infra**: Pasta com os pacotes da camada de infraestrutura (Clean Arch) usados pelo recurso (repository e controller).
- **modules/product/infra/controller**: Pasta com o pacote controller do recurso.
- **modules/product/infra/repository**: Pasta com o pacote repository do recurso.
- **modules/product/usecases**: Pasta com os pacotes da camada UseCase (Clean Arch) do recurso.
- **modules/product/usecases/create**: Pasta com o pacote **create**.
- **modules/product/usecases/delete**: Pasta com o pacote **delete**.
- **modules/product/usecases/get**: Pasta com o pacote **get**.
- **modules/product/usecases/get-all**: Pasta com o pacote **get-all**.
- **modules/product/usecases/update**: Pasta com o pacote **update**.

<h3>Fluxo de Execução do endpoint GET /api/products/{id} <span style="font-size: 0.7em;">(Os demais endpoints seguem a mesma lógica de execução)</span></h3>

Considere que a API está em execução na porta **8888 (http://localhost:8888/api/products/034ab7b3-90ea-11f0-95f2-00155d6d5ec0)**
1. A requisição **GET** chega ao endpoint **/api/products/{id}**. A rota é processada pelo pacote **routes**.
2. O pacote **routes** aciona a função **StartGet**, passando as dependências de **log**, **contexto do Gin** e **repository**.
3. Dentro da função **StartGet** (no pacote **products**), uma instância do **UseCase** é criada com as dependências de **log** e **repository**. Em seguida, a função **Get** do **controller** é chamada, recebendo o **log**, o **contexto do Gin** e o **UseCase** como dependências.
4. No pacote **controller**, a função **Get** extrai os parâmetros do endpoint por meio do **contexto do Gin** e chama a função **Execute** do **UseCase** correspondente.
5. A função **Execute** (do **UseCase**) executa a lógica necessária para atender à requisição. Nesse caso, ela invoca o **repository** para buscar os dados do produto.
6. O fluxo retorna à função **Get** do **controller**, que recebe os dados do **UseCase** e os envia para a interface HTTP do endpoint.

## Autor

Este projeto foi desenvolvido por:

*   [Valdinei Valmir dos Santos](https://github.com/valdinei-santos)
