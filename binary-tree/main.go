package main

import (
	"fmt"
	"./tree"
)

func main() {

	tr := new(tree.Tree)

	tr.Add(9)
	tr.Add(15)
	tr.Add(18)
	tr.Add(12)
	tr.Add(13)
	tr.Add(10)
	tr.Add(11)
	tr.Add(5)
	tr.Add(7)
	tr.Add(6)
	tr.Add(8)
	tr.Add(2)
	tr.Add(1)
	tr.Add(4)

	fmt.Println("\nÁrvore:")
	tr.PrintInOrder()

	fmt.Println("----------------------------------------")
	tr.Remove(12)

	//fmt.Println("\n\nÁrvore:")
	tr.PrintInOrder()
}