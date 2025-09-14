package datafake

import (
	"fmt"
	"math/rand"

	"github.com/valdinei-santos/product-details/modules/product/domain/entities"
	"github.com/valdinei-santos/product-details/modules/product/infra/repository"
)

type Produto struct {
	ID            string  `json:"ID"`
	Nome          string  `json:"Nome"`
	URL           string  `json:"URL"`
	Descricao     string  `json:"Descricao"`
	Preco         float64 `json:"Preco"`
	Classificacao string  `json:"Classificacao"`
	Especificacao string  `json:"Especificacao"`
}

// GerarProdutosFake gera uma lista de produtos com dados falsos.
func GerarProdutosFake(repoProducts repository.IProductRepository, quantidade int) error {
	fmt.Println("Gerando produtos fake caso tenha menos que 5 produtos...")

	count, err := repoProducts.Count()
	if err != nil {
		fmt.Printf("Erro ao contar produtos: %v", err)
		return fmt.Errorf("erro ao contar produtos: %w", err)
	}

	// Só cria produtos fake se não tiver pelo menos 5 produtos
	if count < 5 {
		// Listas de dados pré-definidos
		nomes := []string{"Smart TV", "Notebook", "Smartphone", "Fone de Ouvido", "Console de Videogame", "Tablet"}
		descricoes := []string{"Produto de alta qualidade", "Design moderno e elegante", "Com as melhores tecnologias", "Ideal para o dia a dia", "Excelente desempenho"}
		classificacoes := []string{"eletronicos", "eletrodomesticos", "informatica", "acessorios", "jogos"}

		for i := 0; i < quantidade; i++ {
			// Escolhe um nome, descrição e classificação aleatórios das listas
			nome := nomes[rand.Intn(len(nomes))]
			url := fmt.Sprintf("http://produto.com/imagem%d.jpg", i+1)
			descricao := descricoes[rand.Intn(len(descricoes))]
			classificacao := classificacoes[rand.Intn(len(classificacoes))]
			preco := float64(rand.Intn(100000)) / 100
			especificacao := fmt.Sprintf("COD-%d", rand.Intn(9999999))

			p, err := entities.NewProduct(nome, url, descricao, preco, classificacao, especificacao)
			if err != nil {
				fmt.Printf("Erro ao criar um novo produto fake: %v", err)
				return err
			}

			// Salva o produto fake no repositório
			err = repoProducts.AddProduct(p)
			if err != nil {
				fmt.Printf("Erro ao adicionar produto fake: %v", err)
				return err
			}
		}
	}
	return nil
}
