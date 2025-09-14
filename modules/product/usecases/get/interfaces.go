package get

import "github.com/valdinei-santos/product-details/modules/product/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(id string) (*dto.Response, error)
}
