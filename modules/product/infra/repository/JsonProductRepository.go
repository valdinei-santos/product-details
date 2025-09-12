package repository

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
)

// ProductRepo é um repositório para gerenciar produtos
type ProductRepo struct {
	filePath string
	products []*entities.Product
	mutex    sync.Mutex
}

// NewProductRepo cria uma nova instância do repositório
func NewProductRepo(filePath string) (*ProductRepo, error) {
	repo := &ProductRepo{
		filePath: filePath,
	}
	err := repo.Load()
	if err != nil {
		return nil, err
	}
	return repo, nil
}

// Load carrega os dados do arquivo JSON para o repositório
func (r *ProductRepo) Load() error {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		// Se o arquivo não existe, inicializa a lista vazia
		if os.IsNotExist(err) {
			r.products = []*entities.Product{}
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &r.products)
}

// Save salva os dados do repositório no arquivo JSON
func (r *ProductRepo) Save() error {
	data, err := json.MarshalIndent(r.products, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

// AddProduct adiciona um novo produto ao repositório
func (r *ProductRepo) AddProduct(p *entities.Product) error {
	r.mutex.Lock()         // Bloqueia o acesso ao repositório
	defer r.mutex.Unlock() // Libera o acesso ao repositório
	r.products = append(r.products, p)
	return r.Save()
}

// GetProductByID busca um produto por ID
func (r *ProductRepo) GetProductByID(id int) (*entities.Product, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, errors.New("produto não encontrado")
}

// GetAllProducts retorna todos os produtos
func (r *ProductRepo) GetAllProducts(offset int, limit int) ([]*entities.Product, int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 2. Garante que o offset e o limit não extrapolem o tamanho do array
	total := len(r.products)
	if offset >= total {
		return []*entities.Product{}, total
	}
	end := offset + limit
	if end > total {
		end = total
	}

	return r.products[offset:end], total
}

// UpdateProduct atualiza os dados de um produto existente
func (r *ProductRepo) UpdateProduct(id int, p *entities.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, product := range r.products {
		if product.ID == id {
			r.products[i] = p
			return r.Save()
		}
	}
	return errors.New("produto não encontrado")
}

// DeleteProduct remove um produto do repositório
func (r *ProductRepo) DeleteProduct(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, product := range r.products {
		if product.ID == id {
			// Remove o usuário da slice
			r.products = append(r.products[:i], r.products[i+1:]...)
			return r.Save()
		}
	}
	return errors.New("produto não encontrado")
}
