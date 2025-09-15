package getall_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	getall "github.com/valdinei-santos/product-details/modules/product/usecases/get-all"
)

func TestExecute(t *testing.T) {
	// Cria uma instância do mock de repositório para ser usada em todos os cenários
	mockRepo := repository.NewMockProductRepository()

	tests := []struct {
		name         string
		repo         *repository.MockProductRepository
		logger       *logger.MockILogger
		page         int
		size         int
		expectedResp *dto.ProductsPaginatedResponse
		expectedErr  error
		expectDebug  bool
		expectError  bool
	}{
		{
			name:   "Deve retornar a primeira página de produtos com sucesso",
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			page:   1,
			size:   2,
			expectedResp: &dto.ProductsPaginatedResponse{
				Products: []dto.Response{
					{
						ID:            mockRepo.Products[0].ID.String(),
						Nome:          mockRepo.Products[0].Nome.String(),
						URLImagem:     mockRepo.Products[0].URLImagem.String(),
						Descricao:     mockRepo.Products[0].Descricao.String(),
						Preco:         mockRepo.Products[0].Preco.Float64(),
						Classificacao: mockRepo.Products[0].Classificacao.String(),
						Especificacao: mockRepo.Products[0].Especificacao.String(),
					},
					{
						ID:            mockRepo.Products[1].ID.String(),
						Nome:          mockRepo.Products[1].Nome.String(),
						URLImagem:     mockRepo.Products[1].URLImagem.String(),
						Descricao:     mockRepo.Products[1].Descricao.String(),
						Preco:         mockRepo.Products[1].Preco.Float64(),
						Classificacao: mockRepo.Products[1].Classificacao.String(),
						Especificacao: mockRepo.Products[1].Especificacao.String(),
					},
				},
				TotalItems:   len(mockRepo.Products),
				TotalPages:   2, // 3 produtos com 2 por página = 2 páginas
				CurrentPage:  1,
				ItemsPerPage: 2,
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:   "Deve retornar a segunda página de produtos com sucesso",
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			page:   2,
			size:   2,
			expectedResp: &dto.ProductsPaginatedResponse{
				Products: []dto.Response{
					{
						ID:            mockRepo.Products[2].ID.String(),
						Nome:          mockRepo.Products[2].Nome.String(),
						URLImagem:     mockRepo.Products[2].URLImagem.String(),
						Descricao:     mockRepo.Products[2].Descricao.String(),
						Preco:         mockRepo.Products[2].Preco.Float64(),
						Classificacao: mockRepo.Products[2].Classificacao.String(),
						Especificacao: mockRepo.Products[2].Especificacao.String(),
					},
				},
				TotalItems:   len(mockRepo.Products),
				TotalPages:   2,
				CurrentPage:  2,
				ItemsPerPage: 2,
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:   "Deve retornar página vazia para página que não existe",
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			page:   10,
			size:   2,
			expectedResp: &dto.ProductsPaginatedResponse{
				Products:     []dto.Response{},
				TotalItems:   len(mockRepo.Products),
				TotalPages:   2,
				CurrentPage:  10,
				ItemsPerPage: 2,
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:   "Deve retornar todos os produtos quando o limite é maior que o total",
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			page:   1,
			size:   10, // Maior que o número de produtos
			expectedResp: &dto.ProductsPaginatedResponse{
				Products: []dto.Response{
					{
						ID:            mockRepo.Products[0].ID.String(),
						Nome:          mockRepo.Products[0].Nome.String(),
						URLImagem:     mockRepo.Products[0].URLImagem.String(),
						Descricao:     mockRepo.Products[0].Descricao.String(),
						Preco:         mockRepo.Products[0].Preco.Float64(),
						Classificacao: mockRepo.Products[0].Classificacao.String(),
						Especificacao: mockRepo.Products[0].Especificacao.String(),
					},
					{
						ID:            mockRepo.Products[1].ID.String(),
						Nome:          mockRepo.Products[1].Nome.String(),
						URLImagem:     mockRepo.Products[1].URLImagem.String(),
						Descricao:     mockRepo.Products[1].Descricao.String(),
						Preco:         mockRepo.Products[1].Preco.Float64(),
						Classificacao: mockRepo.Products[1].Classificacao.String(),
						Especificacao: mockRepo.Products[1].Especificacao.String(),
					},
					{
						ID:            mockRepo.Products[2].ID.String(),
						Nome:          mockRepo.Products[2].Nome.String(),
						URLImagem:     mockRepo.Products[2].URLImagem.String(),
						Descricao:     mockRepo.Products[2].Descricao.String(),
						Preco:         mockRepo.Products[2].Preco.Float64(),
						Classificacao: mockRepo.Products[2].Classificacao.String(),
						Especificacao: mockRepo.Products[2].Especificacao.String(),
					},
				},
				TotalItems:   len(mockRepo.Products),
				TotalPages:   1,
				CurrentPage:  1,
				ItemsPerPage: 10,
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name: "Deve retornar erro quando o repositório falha",
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(errors.New("erro de conexão com o banco de dados"))
				return r
			}(),
			logger:       logger.NewMockILogger(),
			page:         1,
			size:         2,
			expectedResp: nil,
			expectedErr:  errors.New("erro de conexão com o banco de dados"),
			expectDebug:  true,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := getall.NewUseCase(tt.repo, tt.logger)

			resp, err := uc.Execute(tt.page, tt.size)

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectDebug, tt.logger.DebugCalled)
			assert.Equal(t, tt.expectError, tt.logger.ErrorCalled)
		})
	}
}
