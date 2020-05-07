package main

import (
	"fmt"
	"time"
)

func create_ticker(d time.Duration) chan struct{} {
	ticker := make(chan struct{})
	go func() {
		for {
			time.Sleep(d)
			ticker <- struct{}{}
		}
	}()
	return ticker
}

func main() {

	ticker := create_ticker(1 * time.Second)

	sentinel := true

	for range ticker {
		if sentinel {
			fmt.Println("tic")
		} else {
			fmt.Println("tac")
		}
		sentinel = !sentinel
	}

	// Note que a goroutine executada dentro de create_ticker será executada para sempre, gerando uma goroutine leak.
	// Caso o bloco do for seja encerrado, essa goroutine continuará executando

}