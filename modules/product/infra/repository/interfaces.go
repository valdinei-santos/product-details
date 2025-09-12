package repository

import "github.com/valdinei-santos/product-details/modules/product/domain/entities"

// IProductRepository define a interface para as operações de product
type IProductRepository interface {
	AddProduct(p *entities.Product) error
	GetProductByID(id int) (*entities.Product, error)
	GetAllProducts(offset int, limit int) ([]*entities.Product, int)
	UpdateProduct(id int, p *entities.Product) error
	DeleteProduct(id int) error
}
