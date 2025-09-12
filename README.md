# API product-details

## Descrição Técnica

A API **product-details** foi criada com o intuito de demonstrarmos uma simples implementação de API Rest usando os conceitos de Arquitetura Limpa e DDD com Golang.
Para facilitar os testes o repositório usado é um arquivo em JSON, podendo outros tipos de repositórios serem adicionados posteriormente de forma totalmente desacoplada.

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

2. Compile o programa:
   ```bash
   make build
   ```
   Saída:
   ```bash
   Build completo: Executável da API product-details gerado
   ```

3. Execute a API rodando o arquivo gerado no Build:
   ```bash
   ./product-details
   ```

   Saída:
   ```bash
   Iniciando...
   Iniciou Log...
   Iniciou Database...
   {"time":"2025-09-12T06:19:08.180555078-03:00","level":"INFO","msg":"start product-details","PORT:":"8888"}
   ```

## Instruções para Rodar os Testes

### Passos

1. Navegue até o diretório do projeto:
   ```bash
   cd product-details
   ```

2. Execute todos os testes:
   ```bash
   make test
   ```

   Exemplo de saída:
   ```bash
   
   ```

3. Para rodar um teste específico, use o comando **go test**, conforme abaixo:
   ```bash
   go test -run TestNomeDoTeste ./caminho/do/pacote
   ```

4. Para ver a cobertura dos testes na aplicação:
   ```bash
   make cover
   ``` 


### Estrutura de Testes
O projeto inclui testes automatizados para os seguintes componentes:

- **abc**: Valida os argumentos de entrada.
- **abc2**: Testa 
- **abc3**: Testa 
- **main**: Testa o fluxo completo do programa, incluindo cenários de sucesso e falha.


## Detalhes de Implementação

### Estrutura do Projeto

- **cmd/api/**: Contém o ponto de entrada principal da API (`main.go`). Onde o programa inicia as estruturas básicas de config, log, database, framework [GIN](https://gin-gonic.com/) e alocada a porta informada no .env para uso da API.
- **cmd/api/routes**: Pacote com a definição dos **endpoints** disponíveis pela API.
- **infra**: Pasta com os pacotes da camada de infraestrutura usados no contexto geral da API (config, database e log).
- **modules**: Pasta com todos os recursos da API, que nesse caso é apenas **product**.
- **modules/product**: Pasta principal do recurso **product**.
- **modules/product/domain**: Pasta da camada dominio da Clean Arch.
- **modules/product/domain/entities**: Pasta com o pacote das entidades do recurso.
- **modules/product/domain/vo**: Pasta com o pacote dos Value Objects definidos.
- **modules/product/dto**: Pasta com o pacote de DTOs definidos para o recurso.
- **modules/product/infra**: Pasta com os pacotes da camada de infraestrutura usados pelo recurso (repository e controller).
- **modules/product/infra/controller**: Pasta com o pacote controller do recurso.
- **modules/product/infra/repository**: Pasta com o pacote repository do recurso.
- **modules/product/usecases**: Pasta com os pacotes da camada UseCase do recurso.
- **modules/product/usecases/create**: Pasta com o pacote **create**.
- **modules/product/usecases/delete**: Pasta com o pacote **delete**.
- **modules/product/usecases/get**: Pasta com o pacote **get**.
- **modules/product/usecases/get-all**: Pasta com o pacote **get-all**.
- **modules/product/usecases/update**: Pasta com o pacote **update**.

<h3>Fluxo de Execução do endpoint GET /api/products/{id} <span style="font-size: 0.7em;">(Os demais endpoints seguem a mesma lógica de execução)</span></h3>

Considere que a API está em execução na porta **8888 (http://localhost:8888/api/products/1)**
1. A requisição **GET** chega ao endpoint **/api/products/{id}**. A rota é processada pelo pacote **routes**.
2. O pacote **routes** aciona a função **StartGet**, passando as dependências de **log**, **contexto do Gin** e **repository**.
3. Dentro da função **StartGet** (no pacote **products**), uma instância do **UseCase** é criada com as dependências de **log** e **repository**. Em seguida, a função **Get** do **controller** é chamada, recebendo o **log**, o **contexto do Gin** e o **UseCase** como dependências.
4. No pacote **controller**, a função **Get** extrai os parâmetros do endpoint por meio do **contexto do Gin** e chama a função **Execute** do **UseCase** correspondente.
5. A função **Execute** (do **UseCase**) executa a lógica necessária para atender à requisição. Nesse caso, ela invoca o **repository** para buscar os dados do produto.
6. O fluxo retorna à função **Get** do **controller**, que recebe os dados do **UseCase** e os envia para a interface HTTP do endpoint.

## Autor

Este projeto foi desenvolvido por:

*   [Valdinei Valmir dos Santos](https://github.com/valdinei-santos)
