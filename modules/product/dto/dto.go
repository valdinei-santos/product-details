package dto

// Request - Representa os dados necessários para criar ou atualizar um produto
type Request struct {
	ID            string  `json:"id"`
	Nome          string  `json:"nome"`
	URLImagem     string  `json:"url_imagem"`
	Descricao     string  `json:"descricao"`
	Preco         float64 `json:"preco"`
	Classificacao string  `json:"classificacao"`
	Especificacao string  `json:"especificacao"`
}

// Response - Representa a resposta de um produto único
type Response struct {
	ID            string  `json:"id"`
	Nome          string  `json:"nome"`
	URLImagem     string  `json:"url_imagem"`
	Descricao     string  `json:"descricao"`
	Preco         float64 `json:"preco"`
	Classificacao string  `json:"classificacao"`
	Especificacao string  `json:"especificacao"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type ResponseManyPaginated struct {
	Products     []Response `json:"products"`
	TotalItems   int        `json:"totalItems"`
	TotalPages   int        `json:"totalPages"`
	CurrentPage  int        `json:"currentPage"`
	ItemsPerPage int        `json:"itemsPerPage"`
}

// ResponseMany - Representa a resposta de uma lista de produtos
type ResponseMany struct {
	Products []Response `json:"products"`
}

// OutputDefault - Struct com a resposta padrão da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
