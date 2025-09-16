package compare_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/infra/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/compare"
)

func TestExecute(t *testing.T) {
	mockRepo := repository.NewMockProductRepository()
	// Pega os IDs dos produtos padrão do mock
	ids := []string{mockRepo.Products[0].ID.String(), mockRepo.Products[1].ID.String()}
	fmt.Println(ids)

	// Tabela de teste para cenários
	tests := []struct {
		name         string
		repo         *repository.MockProductRepository
		logger       *logger.MockILogger
		inputIDs     []string
		expectedResp *dto.ResponseMany
		expectedErr  error
		expectDebug  bool
		expectError  bool
	}{
		{
			name:     "Deve devolver os produtos em caso de sucesso",
			repo:     mockRepo,
			logger:   logger.NewMockILogger(),
			inputIDs: ids,
			expectedResp: &dto.ResponseMany{
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
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name:         "Deve retornar erro quando repositório falhar",
			repo:         repository.NewMockProductRepository(), // Nova instância para evitar efeitos colaterais
			logger:       logger.NewMockILogger(),
			inputIDs:     []string{"id-invalido"},
			expectedResp: nil,
			expectedErr:  localerror.ErrProductIDInvalid, //errors.New("ID inválido: invalid UUID length: 11"),
			expectDebug:  true,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepara o UseCase com os mocks
			uc := compare.NewUseCase(tt.repo, tt.logger)

			// Executa a função e verifica os resultados
			resp, err := uc.Execute(tt.inputIDs)

			// Asserts
			assert.Equal(t, tt.expectedResp, resp)
			// Verifica se o erro é o esperado usando assert.EqualError
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
