package repository

import (
	"encoding/json"

	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
	"github.com/valdinei-santos/product-details/modules/product/domain/localerror"
)

// ProductRepo é um repositório para gerenciar produtos
type ProductRepo struct {
	filePath string
	products []*entities.Product
	mutex    sync.Mutex
}

// NewProductRepo - cria uma nova instância do repositório
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

// Load - carrega os dados do arquivo JSON para o repositório
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

// Save - salva os dados do repositório no arquivo JSON
func (r *ProductRepo) Save() error {
	data, err := json.MarshalIndent(r.products, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

// AddProduct - adiciona um novo produto ao repositório
func (r *ProductRepo) AddProduct(p *entities.Product) error {
	r.mutex.Lock()         // Bloqueia o acesso ao repositório/Arquivo JSON
	defer r.mutex.Unlock() // Libera o acesso ao repositório/Arquivo JSON
	r.products = append(r.products, p)
	return r.Save()
}

// GetProductByID - busca um produto por ID
func (r *ProductRepo) GetProductByID(id string) (*entities.Product, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, localerror.ErrProductIDInvalid
	}
	for _, product := range r.products {
		if product.ID == idUUID {
			return product, nil
		}
	}
	return nil, localerror.ErrProductNotFound
}

// GetManyProductByIDs - busca vários produtos por ID
func (r *ProductRepo) GetManyProductByIDs(ids []string) ([]*entities.Product, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var products []*entities.Product
	for _, id := range ids {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return nil, localerror.ErrProductIDInvalid
		}
		for _, product := range r.products {
			if product.ID == idUUID {
				products = append(products, product)
			}
		}
	}
	if len(products) == 0 {
		return nil, localerror.ErrProductNotFoundMany
	}
	return products, nil
}

// GetAllProducts - retorna todos os produtos
func (r *ProductRepo) GetAllProducts(offset int, limit int) ([]*entities.Product, int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Garante que o offset e o limit não extrapolem o tamanho do array
	total := len(r.products)
	if offset >= total {
		return []*entities.Product{}, total, nil
	}
	end := offset + limit
	if end > total {
		end = total
	}

	return r.products[offset:end], total, nil
}

// UpdateProduct - atualiza os dados de um produto existente
func (r *ProductRepo) UpdateProduct(id string, p *entities.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return localerror.ErrProductIDInvalid
	}
	for i, product := range r.products {
		if product.ID == idUUID {
			r.products[i] = p
			return r.Save()
		}
	}
	return localerror.ErrProductNotFound
}

// DeleteProduct - remove um produto do repositório
func (r *ProductRepo) DeleteProduct(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return localerror.ErrProductIDInvalid
	}
	for i, product := range r.products {
		if product.ID == idUUID {
			// Remove o produto da slice
			r.products = append(r.products[:i], r.products[i+1:]...)
			return r.Save()
		}
	}
	return localerror.ErrProductNotFound
}

func (r *ProductRepo) Count() (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return len(r.products), nil
}
