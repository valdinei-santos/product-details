package compare

import (
	"github.com/valdinei-santos/product-details/modules/product/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(ids []string) (*dto.ResponseMany, error)
}
