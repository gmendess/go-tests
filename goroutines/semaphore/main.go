package main

// Para evitar que muitas goroutines sejam criadas, é importante usar o conceito de semáforos, que
// usa um buffered channel para limitar a quantidade de goroutines criadas. Cada gorountine criada
// consome uma posição do channel, que é liberada após a finalização da goroutine.

import (
	"fmt"
	"time"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // cria semáforo de 10 posições

	for x := 0; x < 1000; x++ {
		semaphore <- struct{}{} // adquire uma posição do channel
		wg.Add(1)

		go func(number int) { 
			fmt.Println(number)
			time.Sleep(200 * time.Millisecond) // sleep de 200 milissegundos para simular tarefa demorada
			wg.Done()	
			<- semaphore // libera uma posição do channel
		}(x) // goroutine

	}

	// observe que os números serão impressoes de 10 em 10, ou seja, 10 goroutines executando por vez. Sem o semáforo, seriam criadas
	// 1000 goroutines e dependendo da tarefa sendo executada por cada uma, o programa poderia apresentar quedas de performance, devido
	// a quantidade de trocas de contexto necessárias para executar todas as 1000 goroutines concorrentemente. 

	wg.Wait() // espera todas as goroutines terminarem

}