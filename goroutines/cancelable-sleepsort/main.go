package main

import (
	"context"
	"time"
	"fmt"
	"os"
	"log"
)

func c_sleepsort(numbers []int, ctx context.Context) []int {
	channel := make(chan int)
	sorted_slice := make([]int, 0, len(numbers))

	// percorre o slice criando as goroutines
	for _, num := range numbers {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			log.Println("-", n)
			channel <- n // envia o número para o channel para ser adicionado ao sorted_slice
		}(num)
	}

	for range numbers {
		select {
			case <-ctx.Done(): // caso a função cancel seja executada, o channel interno do context é fechado, fazendo com q o case capture o valor padrão do channel
				return nil
			case num := <-channel: // número enviado ao channel, adiciona ele em sorted_slice
				sorted_slice = append(sorted_slice, num)
		}
	}

	return sorted_slice
}

func main() {

	ctx := context.Background() // cria um contexto vazio
	ctx, cancel := context.WithCancel(ctx) // a partir do contexto vazio, cria um contexto que pode ser cancelável

	fmt.Println("ENTER to cancel!")
	go func() {
		os.Stdin.Read(make([]byte, 1)) // lê um byte da entrada padrão
		cancel() // caso o byte seja lido, cancela o contexto, que encerra prematuramente a função c_sleepsort
	}()

	numbers := []int{10, 2, 7, 4, 3, 8, 9, 1, 5, 0, 6}
	sorted_slice := c_sleepsort(numbers, ctx)

	if sorted_slice == nil {
		fmt.Println("sleepsort has been canceled!")
	} else {
		fmt.Printf("\noriginal slice: %v\nsorted slice: %v\n", numbers, sorted_slice)
	}

}