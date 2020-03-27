package main

import (
	"fmt"
	"./tree"
)

func main() {

	tr := new(tree.Tree)

	if tr.Add(10) {
		fmt.Println("Nó inserido com sucesso!")
	}

	if tr.Add(8) {
		fmt.Println("Nó inserido com sucesso!")
	}	

	if tr.Add(14) {
		fmt.Println("Nó inserido com sucesso!")
	}

	if tr.Add(14) {
		fmt.Println("Nó inserido com sucesso!")
	} else {
		fmt.Println("Nó repetido!")
	}

	if tr.Search(8) {
		fmt.Println("Nó encontrado!")
	} else {
		fmt.Println("Nó não encontrado!")
	}


	tr.Add(144)
	tr.Add(9)
	tr.Add(5)
	tr.Add(6)
	tr.Add(120)
	tr.Add(130)
	tr.Add(140)
	tr.Add(115)
	tr.Add(150)

	fmt.Println("\nÁrvore:")
	tr.PrintInOrder()
}