package get_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/infra/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/get"
)

func TestExecute(t *testing.T) {
	// Pega um ID válido do mock de repositório
	mockRepoWithProduct := repository.NewMockProductRepository()
	validID := mockRepoWithProduct.Products[0].ID.String()

	tests := []struct {
		name         string
		repo         *repository.MockProductRepository
		logger       *logger.MockILogger
		inputID      string
		expectedResp *dto.Response
		expectedErr  error
		expectDebug  bool
		expectError  bool
	}{
		{
			name:    "Deve retornar sucesso quando o produto é encontrado",
			repo:    mockRepoWithProduct, // Usa o mock que tem o produto
			logger:  logger.NewMockILogger(),
			inputID: validID,
			expectedResp: &dto.Response{
				ID:            mockRepoWithProduct.Products[0].ID.String(),
				Nome:          mockRepoWithProduct.Products[0].Nome.String(),
				URLImagem:     mockRepoWithProduct.Products[0].URLImagem.String(),
				Descricao:     mockRepoWithProduct.Products[0].Descricao.String(),
				Preco:         mockRepoWithProduct.Products[0].Preco.Float64(),
				Classificacao: mockRepoWithProduct.Products[0].Classificacao.String(),
				Especificacao: mockRepoWithProduct.Products[0].Especificacao.String(),
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:         "Deve retornar erro quando o ID não é encontrado",
			repo:         repository.NewMockProductRepository(), // Usa uma nova instância sem o produto
			logger:       logger.NewMockILogger(),
			inputID:      uuid.New().String(), // Gera um ID que não existe
			expectedResp: nil,
			expectedErr:  localerror.ErrProductNotFound,
			expectDebug:  true,
			expectError:  true,
		},
		{
			name:         "Deve retornar erro quando o ID é inválido",
			repo:         repository.NewMockProductRepository(),
			logger:       logger.NewMockILogger(),
			inputID:      "id-invalido", // ID que não é um UUID
			expectedResp: nil,
			expectedErr:  localerror.ErrProductIDInvalid, //errors.New("ID inválido: invalid UUID length: 11"),
			expectDebug:  true,
			expectError:  true,
		},
		{
			name: "Deve retornar erro quando o repositório falha",
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(errors.New("erro de conexão com o banco de dados"))
				return r
			}(),
			logger:       logger.NewMockILogger(),
			inputID:      validID,
			expectedResp: nil,
			expectedErr:  errors.New("erro de conexão com o banco de dados"),
			expectDebug:  true,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := get.NewUseCase(tt.repo, tt.logger)

			resp, err := uc.Execute(tt.inputID)

			assert.Equal(t, tt.expectedResp, resp)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expectDebug, tt.logger.DebugCalled)
			assert.Equal(t, tt.expectError, tt.logger.ErrorCalled)
		})
	}
}
