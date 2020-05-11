package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
)

type Client struct {
	conn net.Conn
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

	return &Client{conn}
}

func (s *Server) HandleClient(c *Client) {
	s.clients[c] = true

	log.Printf("%v entrou!", c.conn.RemoteAddr())
	c.conn.Write([]byte("Boas-vindas!\n"))

	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		s.messages <- fmt.Sprintf("%v: %s\n", c.conn.RemoteAddr(), scanner.Text())
	}

	log.Printf("%v saiu!", c.conn.RemoteAddr())
	c.conn.Close()
}

func (s *Server) Broadcast() {
	for msg := range s.messages {
		for client := range s.clients {
			client.conn.Write([]byte(msg))
		}
	}
}

func (c Client) String() string {
	return fmt.Sprintf("%v", c.conn.RemoteAddr())
}

func main() {
	server := NewTCPServer(":8080")

	go server.Broadcast()
	for {
		client := server.WaitConnection()
		if client == nil {
			continue
		}

		go server.HandleClient(client)
	}
}