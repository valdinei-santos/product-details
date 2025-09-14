package routes_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/cmd/api/routes"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

func TestInitRoutes(t *testing.T) {
	// Mocks
	mockILogger := logger.NewMockILogger()
	mockRepo := repository.NewMockProductRepository()
	validID := mockRepo.Products[0].ID.String()

	// Cria um novo Gin engine e grupo de roteamento
	gin.SetMode(gin.TestMode)
	router := gin.New()
	apiGroup := router.Group("/")

	// Inicializa rotas
	routes.InitRoutes(apiGroup, mockILogger, mockRepo)

	// Casos de teste
	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{"Ping Route", "GET", "/ping", http.StatusOK},
		{"Get All Products", "GET", "/api/products/", http.StatusOK},
		{"Create Product", "POST", "/api/products/", http.StatusOK},
		{"Get Product by ID", "GET", "/api/products/" + validID, http.StatusOK},
		{"Update Product by ID", "PUT", "/api/products/" + validID, http.StatusOK},
		{"Delete Product by ID", "DELETE", "/api/products/" + validID, http.StatusOK},
	}

	// Executa o casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Define o contexto para o caso de teste
			mockILogger.SetContext(tc.name)

			// Define o corpo da requisição para POST e PUT
			var body string
			if tc.method == "POST" || tc.method == "PUT" {
				//body = `{"name": "Test Product", "price": 10.0}`
				body = `{
					"Nome": "Default Product1", 
					"URL": "http://empresa.com/imagem1", 
					"Descricao": "Produto de Teste1", 
					"Preco": 1.0, 
					"Classificacao": "Eletronicos", 
					"Especificacao": "Teste de Especificacao"
				}`
			}

			// Simula a requisição
			req := httptest.NewRequest(tc.method, tc.url, strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json") // Define o cabeçalho como JSON
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			// Verifica o status da resposta
			assert.Equal(t, tc.statusCode, resp.Code)

			// Exibe os logs apenas se o teste falhar
			if t.Failed() {
				logs := mockILogger.GetLogs(tc.name)
				t.Logf("Logs gerados no teste '%s':\n%s", tc.name, strings.Join(logs, "\n"))
			}
		})
	}
}
