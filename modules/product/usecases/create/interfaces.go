package create

import "github.com/valdinei-santos/product-details/modules/product/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(p *dto.Request) (*dto.OutputDefault, error)
}
