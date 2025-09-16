package create

import "github.com/valdinei-santos/product-details/modules/product/infra/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(p *dto.Request) (*dto.Response, error)
}
