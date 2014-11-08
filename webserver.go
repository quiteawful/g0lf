package main

import (
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

type SocketServer struct {
	name    string
	clients map[int]*Client
}

func NewSocketServer(name string) *SocketServer {
	return &SocketServer{
		name,
		make(map[int]*Client),
	}
}

func (s *SocketServer) listen() {
	log.Println("Starting Listener")
	http.Handle(s.name, websocket.Handler(s.wsHandler))
}

func (s *SocketServer) wsHandler(ws *websocket.Conn) {
	log.Println("New Connection")
	c := NewClient(ws)
	s.Add(c)
	c.listen()

}

func (s *SocketServer) Add(c *Client) {
	s.clients[c.id] = c
	log.Println("Added new Client, id: ", c.id)

}
