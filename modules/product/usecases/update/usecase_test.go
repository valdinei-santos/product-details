package update_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
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
		expectedResp *dto.Response
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
				URLImagem:     "http://empresa.com/novo-url",
				Descricao:     "Descrição Atualizada",
				Preco:         50.0,
				Classificacao: "Moda",
				Especificacao: "Nova especificação",
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
				URLImagem:     "http://empresa.com/novo-url",
				Descricao:     "Descrição Atualizada",
				Preco:         50.0,
				Classificacao: "Moda",
				Especificacao: "Nova especificação",
			},
			expectedErr: localerror.ErrProductNotFound,
			expectDebug: true,
			expectError: true,
		},
		{
			name: "Deve retornar erro quando o repositório falha ao atualizar",
			id:   validProductID,
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(localerror.ErrProductNotFound)
				return r
			}(),
			logger: logger.NewMockILogger(),
			input: &dto.Request{
				Nome:          "Produto Atualizado",
				URLImagem:     "http://empresa.com/novo-url",
				Descricao:     "Descrição Atualizada",
				Preco:         50.0,
				Classificacao: "Moda",
				Especificacao: "Nova especificação",
			},
			expectedErr: localerror.ErrProductNotFound,
			expectDebug: true,
			expectError: true,
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
			expectedErr: localerror.ErrProductNameInvalid,
			expectDebug: true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := update.NewUseCase(tt.repo, tt.logger)

			resp, err := uc.Execute(tt.id, tt.input)

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
