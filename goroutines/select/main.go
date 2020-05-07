package main

// A diretiva 'select' espera até que alguma comunicação ocorra em um de seus 'cases'. Dessa forma,
// o case que realizar primeiro uma ação em um channel(enviar ou receber) terá seu bloco executado 

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	carro1 := make(chan int)
	carro2 := make(chan int)

	// goroutine do carro2
	go func() {
		random_time := rand.Intn(2) + 2 
		time.Sleep(time.Duration(random_time) * time.Second)
		carro1 <- random_time
	}()

	// goroutine do carro2
	go func() {
		random_time := rand.Intn(2) + 2 
		time.Sleep(time.Duration(random_time) * time.Second)
		carro2 <- random_time
	}()

	// select espera até que um dos cases receba uma informação vinda ou do canal 'carro1'
	// ou do canal 'carro2'. O canal que receber uma informação primeiro, terá seu respectivo
	// case executado primeiro
	select {
	case t := <-carro1:
		fmt.Printf("Carro 1 terminou primeiro após %d segundos!", t)
	case t := <-carro2:
		fmt.Printf("Carro 2 terminou primeiro após %d segundos!", t)
	}

	// Um select pode ser usado dentro de um loop para criar timeouts, por exemplo, vc cria 2 canais, o do timeout e outro que
	// recebe mensagens. Se o canal do timeout receber uma informação primeiro que o de mensagens, significa que o usuário
	// ficou x segundos sem enviar nada
}