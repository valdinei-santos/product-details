package getall

import "github.com/valdinei-santos/product-details/modules/product/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(page int, size int) (*dto.ResponseManyPaginated, error)
}
