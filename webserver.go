package main

import (
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

type SocketServer struct {
	name    string
	clients map[int]*Client
	addCh   chan *Client
}

func NewSocketServer(name string) *SocketServer {
	return &SocketServer{
		name,
		make(map[int]*Client),
		make(chan *Client),
	}
}

func (s *SocketServer) listen() {
	log.Println("Starting Listener")

	http.Handle(s.name, websocket.Handler(s.wsHandler))
	log.Println("Created handler")
	for {
		select {

		case c := <-s.addCh:
			s.clients[c.id] = c
			log.Println("Added new Client, id: ", c.id)
		}

	}
}

func (s *SocketServer) wsHandler(ws *websocket.Conn) {

	log.Println("New Connection")

	defer func() {
		log.Println("Client closed connection")
		err := ws.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	c := NewClient(ws, s)
	s.Add(c)
	c.listen()
}

func (s *SocketServer) Add(c *Client) {
	s.addCh <- c

}

func (s *SocketServer) Del(c *Client) {
	delete(s.clients, c.id)
	log.Println("Client deleted")
}
