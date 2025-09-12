package repository

import (
	"fmt"

	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
)

// MockProductRepository é um mock com a implementação da interface IProductRepository
type MockProductRepository struct {
	products []entities.Product
}

// NewMockProductRepository cria uma nova instancia de MockProductRepository com 3 produtos padrão
func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: []entities.Product{
			{ID: 1, Nome: "Default Product1", URL: "http://empresa.com/imagem1", Descricao: "Produto de Teste1", Preco: 1.0, Classificacao: "Eletronicos", Especificacao: "Teste"},
			{ID: 2, Nome: "Default Product2", URL: "http://empresa.com/imagem2", Descricao: "Produto de Teste2", Preco: 2.0, Classificacao: "Eletronicos", Especificacao: "Teste"},
			{ID: 3, Nome: "Default Product3", URL: "http://empresa.com/imagem3", Descricao: "Produto de Teste3", Preco: 3.0, Classificacao: "Eletronicos", Especificacao: "Teste"},
		},
	}
}

// GetAllProducts - mock do método GetAllProducts
func (m *MockProductRepository) GetAllProducts(offset int, limit int) ([]*entities.Product, int) {
	total := len(m.products)

	// Aplica o offset e o limit para simular paginação
	if offset > total {
		return []*entities.Product{}, total
	}

	end := offset + limit
	if end > total {
		end = total
	}

	// Converte os produtos para um slice de ponteiros
	products := make([]*entities.Product, 0, end-offset)
	for i := offset; i < end; i++ {
		products = append(products, &m.products[i])
	}

	return products, total
}

// GetProductByID - mock do método GetProductByID
func (m *MockProductRepository) GetProductByID(id int) (*entities.Product, error) {
	for _, product := range m.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

// AddProduct - mock do método AddProduct
func (m *MockProductRepository) AddProduct(p *entities.Product) error {
	if p == nil {
		return fmt.Errorf("produto não pode ser nil")
	}
	// Cria um ID simples baseado no tamanho do slice
	p.ID = len(m.products) + 1
	// Adiciona o produto ao slice
	m.products = append(m.products, *p)
	return nil
}

// UpdateProduct - mock do método UpdateProduct
func (m *MockProductRepository) UpdateProduct(id int, p *entities.Product) error {
	for i, product := range m.products {
		if product.ID == id {
			// Atualiza o produto existente com os novos valores
			p.ID = id // Garante que o ID não seja alterado
			m.products[i] = *p
			return nil
		}
	}
	return fmt.Errorf("produto com ID %d não encontrado", id)
}

// DeleteProduct - mock do método DeleteProduct
func (m *MockProductRepository) DeleteProduct(id int) error {
	for i, p := range m.products {
		if p.ID == id {
			m.products = append(m.products[:i], m.products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("produto não encontrado")
}
