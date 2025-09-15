package create_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/create"
)

func TestExecute(t *testing.T) {
	// Pega um ID válido do mock de repositório
	//mockRepoWithProduct := repository.NewMockProductRepository()
	//validID := mockRepoWithProduct.Products[0].ID.String()

	// Tabela de teste para cenários
	tests := []struct {
		name         string
		repo         *repository.MockProductRepository
		logger       *logger.MockILogger
		input        *dto.Request
		expectedResp *dto.Response
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
				URLImagem:     "http://teste.com/img.jpg",
				Descricao:     "Descricao de Teste",
				Preco:         50.0,
				Classificacao: "eletronicos",
				Especificacao: "Teste de especificação",
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:        "Deve retornar error quando dados inválidos são fornecidos",
			repo:        repository.NewMockProductRepository(),
			logger:      logger.NewMockILogger(),
			input:       &dto.Request{Nome: ""}, // Nome vazio causa erro
			expectedErr: localerror.ErrProductNameInvalid,
			expectDebug: true,
			expectError: true,
		},
		{
			name: "Deve retornar error quando repositório falha",
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(localerror.ErrProductInternal)
				return r
			}(),
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Teste 2",
				URLImagem:     "http://teste.com/img2.jpg",
				Descricao:     "Descricao de Teste 2",
				Preco:         75.0,
				Classificacao: "acessorios",
				Especificacao: "Teste de especificação",
			},
			expectedErr: localerror.ErrProductInternal,
			expectDebug: true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := create.NewUseCase(tt.repo, tt.logger)

			resp, err := uc.Execute(tt.input)

			//Verifique se não houve erro
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				//Verifique somente os campos que não mudam
				assert.Equal(t, tt.input.Nome, resp.Nome)
				assert.Equal(t, tt.input.URLImagem, resp.URLImagem)
				assert.Equal(t, tt.input.Descricao, resp.Descricao)
				assert.Equal(t, tt.input.Preco, resp.Preco)
				assert.Equal(t, tt.input.Classificacao, resp.Classificacao)
				assert.Equal(t, tt.input.Especificacao, resp.Especificacao)

				// Verifique se o ID foi gerado
				assert.NotEmpty(t, resp.ID)

				// Verifique se as datas não estão vazias e estão no formato correto
				assert.NotEmpty(t, resp.CreatedAt)
				assert.NotEmpty(t, resp.UpdatedAt)
			}

			assert.Equal(t, tt.expectDebug, tt.logger.DebugCalled)
			assert.Equal(t, tt.expectError, tt.logger.ErrorCalled)
		})
	}
}
