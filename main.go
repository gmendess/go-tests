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

	if tr.Add(13) {
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
}