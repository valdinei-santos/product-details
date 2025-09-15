package delete_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdinei-santos/product-details/infra/logger"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
	"github.com/valdinei-santos/product-details/modules/product/dto"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
	"github.com/valdinei-santos/product-details/modules/product/usecases/delete"
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
		expectedResp *dto.OutputDefault
		expectedErr  error
		expectDebug  bool
		expectError  bool
	}{
		{
			name:    "Deve retornar sucesso ao excluir um produto",
			repo:    mockRepoWithProduct,
			logger:  logger.NewMockILogger(),
			inputID: validID,
			expectedResp: &dto.OutputDefault{
				StatusCode: 1,
				Message:    "Produto deletado com sucesso",
			},
			expectedErr: nil,
			expectDebug: true,
			expectError: false,
		},
		{
			name: "Deve retornar erro se o repositório falhar",
			repo: func() *repository.MockProductRepository {
				r := repository.NewMockProductRepository()
				r.SetMockError(localerror.ErrProductConnectionInDatabase)
				return r
			}(),
			logger:       logger.NewMockILogger(),
			inputID:      validID,
			expectedResp: nil,
			expectedErr:  localerror.ErrProductConnectionInDatabase,
			expectDebug:  true,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := delete.NewUseCase(tt.repo, tt.logger)

			err := uc.Execute(tt.inputID)

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
