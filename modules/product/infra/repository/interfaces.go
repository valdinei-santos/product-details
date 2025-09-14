package repository

import "github.com/valdinei-santos/product-details/modules/product/domain/entities"

// IProductRepository define a interface para as operações de product
type IProductRepository interface {
	AddProduct(p *entities.Product) error
	GetProductByID(id string) (*entities.Product, error)
	GetManyProductByIDs(ids []string) ([]*entities.Product, error)
	GetAllProducts(offset int, limit int) ([]*entities.Product, int, error)
	UpdateProduct(id string, p *entities.Product) error
	DeleteProduct(id string) error
	Count() (int, error)
}
