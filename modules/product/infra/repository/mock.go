package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
)

// MockProductRepository é um mock com a implementação da interface IProductRepository
type MockProductRepository struct {
	Products  []entities.Product
	mockError error
	callCount int
}

// NewMockProductRepository cria uma nova instancia de MockProductRepository com 3 produtos padrão
func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		Products: []entities.Product{
			{ID: uuid.New(), Nome: "Default Product1", URL: "http://empresa.com/imagem1", Descricao: "Produto de Teste1", Preco: 1.0, Classificacao: "Eletronicos", Especificacao: "Teste"},
			{ID: uuid.New(), Nome: "Default Product2", URL: "http://empresa.com/imagem2", Descricao: "Produto de Teste2", Preco: 2.0, Classificacao: "Eletronicos", Especificacao: "Teste"},
			{ID: uuid.New(), Nome: "Default Product3", URL: "http://empresa.com/imagem3", Descricao: "Produto de Teste3", Preco: 3.0, Classificacao: "Eletronicos", Especificacao: "Teste"},
		},
	}
}

func (m *MockProductRepository) SetMockError(err error) {
	m.mockError = err
}

// GetProductByID - mock do método GetProductByID
func (m *MockProductRepository) GetProductByID(id string) (*entities.Product, error) {
	if m.mockError != nil {
		return nil, m.mockError
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %w", err)
	}
	for _, product := range m.Products {
		if product.ID == idUUID {
			return &product, nil
		}
	}
	return nil, errors.New("produto não encontrado")
}

// GetManyProductByIDs - busca vários produtos por ID
func (m *MockProductRepository) GetManyProductByIDs(ids []string) ([]*entities.Product, error) {
	if m.mockError != nil {
		return nil, m.mockError
	}

	var products []*entities.Product
	for _, id := range ids {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("ID inválido: %w", err)
		}
		for _, product := range m.Products {
			if product.ID == idUUID {
				products = append(products, &product)
			}
		}
	}
	return products, nil
}

// GetAllProducts - mock do método GetAllProducts
func (m *MockProductRepository) GetAllProducts(offset int, limit int) ([]*entities.Product, int, error) {
	if m.mockError != nil {
		return nil, 0, m.mockError
	}

	total := len(m.Products)

	// Aplica o offset e o limit para simular paginação
	if offset > total {
		return []*entities.Product{}, total, nil
	}

	end := offset + limit
	if end > total {
		end = total
	}

	// Converte os produtos para um slice de ponteiros
	products := make([]*entities.Product, 0, end-offset)
	for i := offset; i < end; i++ {
		products = append(products, &m.Products[i])
	}

	return products, total, nil
}

// AddProduct - mock do método AddProduct
func (m *MockProductRepository) AddProduct(p *entities.Product) error {
	if m.mockError != nil {
		return m.mockError
	}
	if p == nil {
		return fmt.Errorf("produto não pode ser nil")
	}
	// Cria um UUID
	p.ID = uuid.New()
	// Adiciona o produto ao slice
	m.Products = append(m.Products, *p)
	return nil
}

// UpdateProduct - mock do método UpdateProduct
func (m *MockProductRepository) UpdateProduct(id string, p *entities.Product) error {
	if m.mockError != nil {
		return m.mockError
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}
	for i, product := range m.Products {
		if product.ID == idUUID {
			// Atualiza o produto existente com os novos valores
			p.ID = idUUID // Garante que o ID não seja alterado
			m.Products[i] = *p
			return nil
		}
	}
	return fmt.Errorf("produto com ID %s não encontrado", id)
}

// DeleteProduct - mock do método DeleteProduct
func (m *MockProductRepository) DeleteProduct(id string) error {
	if m.mockError != nil {
		return m.mockError
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}
	for i, p := range m.Products {
		if p.ID == idUUID {
			m.Products = append(m.Products[:i], m.Products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("produto não encontrado")
}

// Count - mock do método Count
func (r *MockProductRepository) Count() (int, error) {
	return len(r.Products), nil
}
