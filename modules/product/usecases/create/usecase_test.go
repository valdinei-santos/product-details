package create_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/create"
)

func TestExecute(t *testing.T) {
	// Tabela de teste para cenários
	tests := []struct {
		name         string
		repo         *repository.MockProductRepository
		logger       *logger.MockILogger
		input        *dto.Request
		expectedResp *dto.OutputDefault
		expectedErr  error
		expectDebug  bool
		expectError  bool
	}{
		{
			name:   "Deve retornar sucesso com dados válidos",
			repo:   repository.NewMockProductRepository(), // Cria uma nova instância para não compartilhar estado. Dava erro
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Teste",
				URL:           "http://teste.com/img.jpg",
				Descricao:     "Descricao de Teste",
				Preco:         50.0,
				Classificacao: "eletronicos",
				Especificacao: "Teste de especificação",
			},
			expectedResp: &dto.OutputDefault{
				StatusCode: 1,
				Message:    "Produto inserido com sucesso",
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:         "Deve retornar error quando dados inválidos são fornecidos",
			repo:         repository.NewMockProductRepository(),
			logger:       logger.NewMockILogger(),
			input:        &dto.Request{Nome: ""}, // Nome vazio causa erro
			expectedResp: nil,
			expectedErr:  errors.New("o nome deve ter entre 3 e 50 caracteres"),
			expectDebug:  true,
			expectError:  true,
		},
		{
			name: "Deve retornar error quando repositório falha",
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(errors.New("erro ao salvar no banco"))
				return r
			}(),
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Teste 2",
				URL:           "http://teste.com/img2.jpg",
				Descricao:     "Descricao de Teste 2",
				Preco:         75.0,
				Classificacao: "acessorios",
				Especificacao: "Teste de especificação",
			},
			expectedResp: nil,
			expectedErr:  errors.New("erro ao salvar no banco"),
			expectDebug:  true,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := create.NewUseCase(tt.repo, tt.logger)

			resp, err := uc.Execute(tt.input)

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
