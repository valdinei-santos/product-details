package update_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/update"
)

func TestExecute(t *testing.T) {
	mockRepo := repository.NewMockProductRepository()
	validProductID := mockRepo.Products[0].ID.String()

	tests := []struct {
		name         string
		id           string
		repo         *repository.MockProductRepository
		logger       *logger.MockILogger
		input        *dto.Request
		expectedResp *dto.OutputDefault
		expectedErr  error
		expectDebug  bool
		expectError  bool
	}{
		{
			name:   "Deve atualizar um produto com sucesso",
			id:     validProductID,
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Atualizado",
				URL:           "http://empresa.com/novo-url",
				Descricao:     "Descrição Atualizada",
				Preco:         50.0,
				Classificacao: "Moda",
				Especificacao: "Nova especificação",
			},
			expectedResp: &dto.OutputDefault{
				StatusCode: 1,
				Message:    "Produto alterado com sucesso",
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:   "Deve retornar erro ao tentar atualizar produto com ID inválido",
			id:     "id-invalido",
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Atualizado",
				URL:           "http://empresa.com/novo-url",
				Descricao:     "Descrição Atualizada",
				Preco:         50.0,
				Classificacao: "Moda",
				Especificacao: "Nova especificação",
			},
			expectedResp: nil,
			expectedErr:  errors.New("ID inválido: invalid UUID length: 11"),
			expectDebug:  true,
			expectError:  true,
		},
		{
			name: "Deve retornar erro quando o repositório falha ao atualizar",
			id:   validProductID,
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(errors.New("erro de conexão com o banco de dados"))
				return r
			}(),
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Atualizado",
				URL:           "http://empresa.com/novo-url",
				Descricao:     "Descrição Atualizada",
				Preco:         50.0,
				Classificacao: "Moda",
				Especificacao: "Nova especificação",
			},
			expectedResp: nil,
			expectedErr:  errors.New("erro de conexão com o banco de dados"),
			expectDebug:  true,
			expectError:  true,
		},
		{
			name:   "Deve retornar erro ao tentar atualizar com dados de entrada inválidos",
			id:     validProductID,
			repo:   mockRepo,
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:  "",    // Nome vazio
				Preco: -10.0, // Preço inválido
			},
			expectedResp: nil,
			expectedErr:  errors.New("o nome deve ter entre 3 e 50 caracteres"),
			expectDebug:  true,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := update.NewUseCase(tt.repo, tt.logger)

			resp, err := uc.Execute(tt.id, tt.input)

			assert.Equal(t, tt.expectedResp, resp)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectDebug, tt.logger.DebugCalled)
			assert.Equal(t, tt.expectError, tt.logger.ErrorCalled)
		})
	}
}
