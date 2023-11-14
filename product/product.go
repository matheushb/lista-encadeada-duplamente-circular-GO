package product

import "fmt"

type Produto struct {
	Nome   string
	Preco  int
	Codigo int
}

func CreateProduct(nome string, preco int, codigo *int) *Produto {
	*codigo++
	return &Produto{
		Nome:   nome,
		Preco:  preco,
		Codigo: *codigo,
	}
}

func PrintProduto(produto Produto) {
	if &produto != nil {
		fmt.Println("Código: ", produto.Codigo)
		fmt.Println("Nome: ", produto.Nome)
		fmt.Println("Preço: ", produto.Preco)
	}
}
