package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"errors"
	"time"
)

type Client struct {
	conn net.Conn
	name string
}

type Server struct {
	listener net.Listener
	clients map[*Client]bool
	messages chan string
}

func NewTCPServer(address string) *Server {
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	clients := make(map[*Client]bool)
	messages := make(chan string)
	return &Server{conn, clients, messages}
}

func (s *Server) WaitConnection() *Client {
	conn, err := s.listener.Accept()
	if err != nil {
		log.Print(err)
		return nil
	}

	return &Client{conn, ""}
}

func request_name(s *bufio.Scanner) (string, error) {
	if s.Scan() != false {
		return s.Text(), nil
	}
	return "", errors.New(fmt.Sprintf("Não foi possível ler nome do cliente! %v", s.Err()))
}

func (s *Server) HandleClient(c *Client) {

	// cria um scanner para ler as mensagens recebidas do cliente
	scanner := bufio.NewScanner(c.conn)
	
	// solicita o nome do cliente
	fmt.Fprintf(c.conn, "Digite seu nome: ")
	var err error
	if c.name, err = request_name(scanner); err != nil {
		log.Println(err)
		return
	}

	s.messages <- fmt.Sprintf("%s entrou!\n", c.name)

	// após receber receber o nome do cliente e reportar aos outros usuários sua entrada, o cliente é
	// adicionado no map de clientes
	s.clients[c] = true

	client_message := make(chan string)

	// goroutine que recebe as mensagens do cliente e envia para o canal cliente_message
	go func() {
		for scanner.Scan() {
			client_message <- fmt.Sprintf("%s: %s\n", c.name, scanner.Text())
		}
	}()

	// código responsável por gerenciar o timeout do cliente
	var ticker *time.Ticker
	loop: 
	for {
		ticker = time.NewTicker(2 * time.Minute) // cada iteração do for cria um novo ticker
		select {
			case <-ticker.C: // se alguma informação for recebida do ticker, significa que se passaram 2 minutos sem o cliente enviar mensagens
				c.conn.Write([]byte("Timeout!\n"))
				break loop
			case msg := <-client_message:
				s.messages <- msg
				ticker.Stop() // para o ticker atual para outro ser criado
		}
	}

	delete(s.clients, c) // deleta o cliente antes de replicar sua saida
	s.messages <- fmt.Sprintf("%s saiu!\n", c.name) // avisa a todos do chat que o cliente se desconectou
	c.conn.Close()
}

func (s *Server) Broadcast() {
	// espera que alguma informação seja enviada para o canal s.messages; a informação recebida é inserida em 'msg'
	for msg := range s.messages {
		// replica para todos os clientes a mensagem recebida
		for client := range s.clients {
			if _, err := client.conn.Write([]byte(msg)); err != nil {
				log.Printf("Erro ao replicar mensagem %q! %v", msg, err)
			}
		}
	}
}

func main() {
	server := NewTCPServer(":8080")

	// goroutine que replica as mensagens recebidas de um cliente
	go server.Broadcast()

	// inicia o loop que espera por conexões
	for {
		client := server.WaitConnection()
		if client == nil {
			// não foi possível abrir conexão com o cliente, espera por outra conexão
			continue
		}

		// lida com a conexão recém criada
		go server.HandleClient(client)
	}
}