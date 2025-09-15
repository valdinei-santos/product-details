# API REST product-details

## Design da API
Esta API é construída com uma arquitetura **RESTful**, usando URLs intuitivas para representar recursos. Ela se baseia nos seguintes princípios:

-   Uso de **verbos HTTP** para descrever as ações sobre os recursos (GET para leitura, POST para criação, PUT para atualização e DELETE para exclusão).
-   Utilização de **códigos de status HTTP** padrão para indicar o resultado da requisição.
-   Todas as requisições e respostas usam o formato **JSON**.

### Endpoints

| Verbo   | Rota                                               | Descrição                          |
| :------ | :------------------------------------------------- | :--------------------------------- |
| `GET`   | `/api/v1/products?page=1&limit=2`                  | Lista todos os produtos.           |
| `POST`  | `/api/v1/products`                                 | Cria um novo produto.              |
| `GET`   | `/api/v1/products{id}`                             | Retorna um produto por ID.         |
| `PUT`   | `/api/v1/products/{id}`                            | Atualiza um produto por ID.        |
| `DELETE`| `/api/v1/products/{id}`                            | Deleta um produto por ID.          |
| `GET   `| `/api/v1/products/compare?ids={id1,id2,id3,...}`   | Deleta um produto por ID.          |


### Tratamento de Erros

Respostas de erro seguem o formato JSON e incluem uma mensagem descritiva:

**Exemplo de erro `404 Not Found`:**

```json
{
  "status_code": -1,
  "message": "produto não encontrado"
}

## Instruções de Configuração

## Decisões Arquiteturais

## Estratégia Técnica

## Como usei IA



## Descrição Técnica

**product-details** é uma API REST em Golang que implementa os princípios de Arquitetura Limpa e DDD para gerenciar informações de produtos.  
A arquitetura foi projetada para garantir o desacoplamento do repositório de dados. A persistência inicial usa um arquivo JSON, mas a estrutura permite a fácil substituição por outros tipos de bancos de dados (como SQL ou NoSQL) sem afetar o código principal da aplicação. A API é populada com dados simulados para facilitar o uso e testes imediatos.


## Documentação da API

Esta API utiliza a especificação OpenAPI (anteriormente conhecida como Swagger) para sua documentação interativa. Para acessá-la, basta iniciar a API e navegar até o seguinte endpoint no seu navegador:

http://localhost:8888/swagger/index.html


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
