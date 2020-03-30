package main

import (
	"fmt"
	"./tree"
)

func main() {

	tr := new(tree.Tree)

	tr.Add(5)	
	tr.Add(7)	
	tr.Add(3)	
	tr.Add(4)	
	tr.Add(2)	

	fmt.Println("\nÁrvore:")
	tr.PrintInOrder()

	fmt.Println("----------------------------------------")
	tr.RotateRight()

	tr.PrintInOrder()
}