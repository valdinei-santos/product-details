# API REST product-details

## Descrição Técnica

**product-details** é uma API REST em Golang que implementa os princípios de Arquitetura Limpa e DDD para gerenciar informações de produtos.  
A arquitetura foi projetada para garantir o desacoplamento do repositório de dados. A persistência inicial usa um arquivo JSON, mas a estrutura permite a fácil substituição por outros tipos de bancos de dados (como SQL ou NoSQL) sem afetar o código principal da aplicação. A API é populada com dados simulados para facilitar o uso e testes imediatos.


## Design da API
Esta API é construída com uma arquitetura **RESTful**, usando URLs intuitivas para representar os recursos. Ela se baseia nos seguintes princípios:

-   Uso de **verbos HTTP** para descrever as ações sobre os recursos (GET para leitura, POST para criação, PUT para atualização e DELETE para exclusão).
-   Utilização de **códigos de status HTTP** padrão para indicar o resultado da requisição.
-   Todas as requisições e respostas usam o formato **JSON**.

### Endpoints

| Verbo   | Rota                                               | Descrição                             |
| :------ | :------------------------------------------------- | :------------------------------------ |
| `GET`   | `/api/v1/products?page=1&limit=2`                  | Lista todos os produtos.              |
| `GET`   | `/api/v1/products{id}`                             | Retorna um produto por ID.            |
| `GET   `| `/api/v1/products/compare?ids={id1,id2,id3,...}`   | Compara uma lista de produtos por ID. |
| `POST`  | `/api/v1/products`                                 | Cria um novo produto.                 |
| `PUT`   | `/api/v1/products/{id}`                            | Atualiza um produto por ID.           |
| `DELETE`| `/api/v1/products/{id}`                            | Deleta um produto por ID.             |


### Tratamento de Erros

Respostas de erro seguem o formato JSON e incluem uma mensagem descritiva:

**Exemplo de erro `404 Not Found`:**

```json
{
  "status_code": -1,
  "message": "produto não encontrado"
}
```

## Decisões Arquiteturais

A arquitetura deste projeto foi pensada para garantir uma maior manutenibilidade e testabilidade, priorizando a separação de responsabilidades e a clareza do código. 
As seguintes decisões foram tomadas para atingir esses objetivos:

**Golang:** Escolhi Golang por ser a linguagem que tenho mais contato atualmente e por viabilizar a construção de APIs de alto desempenho.

**Framework Web GIN:** O GIN foi usado por ser um framework minimalista e de alta performance para o Go. Ele nos permite criar endpoints de API REST de forma rápida e eficiente, fornecendo funcionalidades essenciais como roteamento e middleware, sem impor complexidade desnecessária à arquitetura.

**Arquitetura Limpa (Clean Architecture):** Adotei a Arquitetura Limpa para desacoplar a lógica de negócio das camadas de infraestrutura (como banco de dados). Isso garante que o núcleo da aplicação seja independente de detalhes técnicos, tornando o sistema mais resiliente a mudanças e mais fácil de testar, já que a lógica de domínio pode ser testada isoladamente.

**Conceitos de DDD (Domain-Driven Design):** A aplicação utiliza conceitos de DDD, como Value Objects, para modelar o domínio de forma mais expressiva e robusta. A utilização de Value Objects garante a validade e imutabilidade de certos dados.


## Documentação da API

Esta API utiliza a especificação OpenAPI (anteriormente conhecida como Swagger) para sua documentação interativa. Para acessá-la, basta iniciar a API e navegar até o seguinte endpoint no seu navegador:

http://localhost:8888/swagger/index.html


## Detalhes de Implementação

### Estrutura básica do Projeto

```bash
.
├── cmd
│   └── api
│       ├── docs      # Arquivos do Swagger gerados automaticamente pela ferramenta
│       ├── main.go   # Inicia as estruturas básicas de config, log, database e framework GIN
│       └── routes    # Onde os endpoint estão definidos
├── infra        # Camada de infra usada no contexto geral da API
│   ├── config             # Onde temos as configurações básicas da API
│   ├── database           # Onde fica o arquivo JSON responsável pelo repositório
│   │   ├── datafake       # Onde são gerados dados fake.
│   │   └── products.json  # Arquivo JSON que armazena as informações dos produtos 
│   └── logger             # Onde temos as definições de log da API 
├── Makefile  # Onde definimos algumas automações da API
├── modules   # Onde ficam os recursos da API, nesse caso só o product
│   └── product
│       ├── domain         # Camada de domain do recurso
│       │   ├── entities     # Entidades usadas pelo recurso
│       │   ├── localerror   # Erros usados pelo recurso
│       │   └── vo           # Value Objects usados pelo recurso
│       ├── dto              # DTOs usados pelo recurso
│       ├── infra          # Camada de Infra usada no contexto apenas do recurso
│       │   ├── controller   # Onde temos o Handler/Controller do recurso
│       │   └── repository   # Onde temos o repository do recurso
│       ├── start.go       # Arquivo responsável pela instanciação das dependências que serão usadas por cada usecase
│       └── usecases       # Camada de UseCases do recurso
│           ├── compare      # UseCase compare
│           ├── create       # UseCase create
│           ├── delete       # UseCase delete
│           ├── get          # UseCase get
│           ├── get-all      # UseCase getall
│           └── update       # UseCase update
├── README.md   # Arquivo com informações da API
├── run.md      # Arquivo com informações sobre como instalar/executar a API
└── test.http   # Arquivo usado para testar os endpoints no VSCode com a extensão Client Rest instalada
```

<h3>Fluxo de Execução do endpoint GET /api/products/{id} <span style="font-size: 0.7em;">(Os demais endpoints seguem a mesma lógica de execução)</span></h3>

Com a API em execução na porta **8888**  
**(http://localhost:8888/api/products/034ab7b3-90ea-11f0-95f2-00155d6d5ec0)**
1. A requisição **GET** chega ao endpoint **/api/products/{id}**. A rota é processada pelo pacote **routes**.
2. O pacote **routes** aciona a função **StartGet**, passando as dependências de **log**, **contexto do Gin** e **repository**.
3. Dentro da função **StartGet** (no pacote **products**), uma instância do **UseCase** é criada com as dependências de **log** e **repository**. Em seguida, a função **Get** do **controller** é chamada, recebendo o **log**, o **contexto do Gin** e o **UseCase** como dependências.
4. No pacote **controller**, a função **Get** extrai os parâmetros do endpoint por meio do **contexto do Gin** e chama a função **Execute** do **UseCase** correspondente.
5. A função **Execute** (do **UseCase**) executa a lógica necessária para atender à requisição. Nesse caso, ela invoca o **repository** para buscar os dados do produto.
6. O fluxo retorna à função **Get** do **controller**, que recebe os dados do **UseCase** e os envia para a interface HTTP do endpoint.


## Uso de IA na construção da API
A IA foi usada para auxiliar nas seguintes tarefas:
- Criação dos Testes Automatizados e respectivos Mocks, onde foi necessário fazer ajustes consideráveis para que os testes funcionassem conforme o esperado. 
- Implantação da lib **swaggo**, para fornecer a documentação da API.
- Criação da função geradora de produtos Fake.
- Preenchimento automático de alguns códigos sugeridos pelo Copilot no VSCode.


## Autor

Este projeto foi desenvolvido por:

*   [Valdinei Valmir dos Santos](https://github.com/valdinei-santos)
