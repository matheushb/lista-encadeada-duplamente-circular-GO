package list

import (
	"fmt"
	"main/product"
)

type No struct {
	Produto  *product.Produto
	Proximo  *No
	Anterior *No
}

type Lista struct {
	Inicio  *No
	Fim     *No
	Tamanho int
}

func Insert(lista *Lista, produto *product.Produto) {
	novo := &No{
		Produto:  produto,
		Proximo:  nil,
		Anterior: nil,
	}

	if lista.Inicio == nil {
		lista.Inicio = novo
		lista.Fim = novo
		novo.Proximo = novo
		novo.Anterior = novo
		lista.Tamanho++
		return
	}

	aux := lista.Inicio
	for i := 0; i < lista.Tamanho; i++ {
		if novo.Produto.Nome < aux.Produto.Nome {
			if aux == lista.Inicio {
				novo.Proximo = lista.Inicio
				novo.Anterior = lista.Fim
				lista.Inicio.Anterior = novo
				lista.Inicio = novo
				lista.Fim.Proximo = novo
			} else {
				novo.Proximo = aux
				novo.Anterior = aux.Anterior
				aux.Anterior.Proximo = novo
				aux.Anterior = novo
			}
			lista.Tamanho++
			return
		}
		aux = aux.Proximo
	}

	novo.Anterior = lista.Fim
	novo.Proximo = lista.Inicio
	lista.Fim.Proximo = novo
	lista.Inicio.Anterior = novo
	lista.Fim = novo
	lista.Tamanho++
}
func Remove(lista *Lista, cod int) *product.Produto {
	if lista.Inicio == nil {
		println("Lista vazia.")
		return nil
	}

	aux := lista.Inicio
	for i := 0; i < lista.Tamanho; i++ {
		if aux.Produto.Codigo == cod {
			if aux == lista.Inicio {
				lista.Inicio = lista.Inicio.Proximo
				lista.Inicio.Anterior = lista.Fim
				lista.Fim.Proximo = lista.Inicio
			} else if aux == lista.Fim {
				lista.Fim = lista.Fim.Anterior
				lista.Fim.Proximo = lista.Inicio
				lista.Inicio.Anterior = lista.Fim
			} else {
				aux.Anterior.Proximo = aux.Proximo
				aux.Proximo.Anterior = aux.Anterior
			}
			lista.Tamanho--
			return aux.Produto
		}
		aux = aux.Proximo
	}
	return nil
}

func Print(lista Lista) {
	aux := lista.Inicio
	fmt.Print("\nLista: \n\n")

	for i := 0; i < lista.Tamanho; i++ {
		product.PrintProduto(*aux.Produto)
		aux = aux.Proximo
		fmt.Println()
	}
	fmt.Printf("Tamanho Lista: %d\n", lista.Tamanho)
}

func PrintReverse(lista Lista) {
	aux := lista.Fim
	fmt.Print("\nLista Reversa: \n\n")

	for i := 0; i < lista.Tamanho; i++ {
		product.PrintProduto(*aux.Produto)
		aux = aux.Anterior
		fmt.Println()
	}
	fmt.Printf("Tamanho Lista: %d\n", lista.Tamanho)

}
