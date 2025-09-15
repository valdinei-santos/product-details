package update

import "github.com/valdinei-santos/product-details/modules/product/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(id string, p *dto.Request) (*dto.Response, error)
}
