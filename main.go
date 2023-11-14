package main

import (
	"bufio"
	"fmt"
	"main/list"
	"main/product"
	"os"
)

func main() {
	opt := -1
	nome := ""
	codigo := 0
	preco := 0
	codProduto := 0

	lista := list.Lista{
		Inicio:  nil,
		Fim:     nil,
		Tamanho: 0,
	}

	list.Insert(&lista, product.CreateProduct("Desodorante", 2500, &codigo))
	list.Insert(&lista, product.CreateProduct("Banana", 540, &codigo))
	list.Insert(&lista, product.CreateProduct("Macarrão", 1250, &codigo))
	list.Insert(&lista, product.CreateProduct("Molho de Tomate", 650, &codigo))

	reader := bufio.NewReader(os.Stdin)

	for opt != 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("O que deseja fazer?\n1 - Cadastrar Produto\n2 - Consultar",
			"Produtos\n3 - Consultar Produtos Reverso\n4 - Excluir Produto\n0 - Sair")

		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Println("Digite o nome do produto: ")
			fmt.Scanln(&nome)
			fmt.Println("Digite o preço do produto: ")
			fmt.Scanln(&preco)
			list.Insert(&lista, product.CreateProduct(nome, preco, &codigo))
			fmt.Println("\nProduto adicionado a lista!\n\nPressione ENTER para continuar")
			reader.ReadString('\n')
		case 2:

			list.Print(lista)
			fmt.Println("\n\nPressione ENTER para continuar")
			reader.ReadString('\n')
		case 3:

			list.PrintReverse(lista)
			fmt.Println("\n\nPressione ENTER para continuar")
			reader.ReadString('\n')
		case 4:
      list.Print(lista)

			fmt.Println("Digite o código de um produto para deletar: ")
			fmt.Scanln(&codProduto)
			removedProduto := list.Remove(&lista, codProduto)
			if removedProduto != nil {
				fmt.Println("\nProduto Removido: ")
				product.PrintProduto(*removedProduto)
			} else {
        fmt.Printf("\nProduto não encontrado com id %d", codProduto)
      }

			fmt.Println("\n\nPressione ENTER para continuar")
			reader.ReadString('\n')
		default:
			if opt != 0 {
				fmt.Println("Opcao Invalida")

			}

		}

	}
  fmt.Print("\033[H\033[2J")
  fmt.Print("Sistema Finalizado.")
  
}
