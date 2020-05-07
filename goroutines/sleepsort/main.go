package main

import (
	"fmt"
	"time"
	"sync"
)

func print_number(n int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(n + 50) * time.Millisecond)
	fmt.Println(n)
	wg.Done()
}

func main() {
	numbers := []int{5, 1, 6, 8, 4, 3, 7, 2, 4, 2, 0}
	
	var wg sync.WaitGroup
	for _, n := range numbers {
		wg.Add(1)
		go print_number(n, &wg)
	}

	wg.Wait() // espera por todas as goroutinas terminarem
}