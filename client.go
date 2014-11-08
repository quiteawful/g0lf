package main

import (
	"log"

	"code.google.com/p/go.net/websocket"
)

var gid int = 0

type Client struct {
	id     int
	con    *websocket.Conn
	quitCh chan bool
}

func NewClient(ws *websocket.Conn) *Client {

	gid++
	return &Client{
		gid,
		ws,
		make(chan bool),
	}
}

func (c *Client) listen() {
	c.listenWrite()
}

func (c *Client) listenWrite() {
	log.Println("Starting Client.listenWrite()")
	websocket.JSON.Send(c.con, "Test123")
	log.Print("Sent msg")
}

func (c *Client) listenRead() {

	var test string
	select {
	case <-c.quitCh:
		//tbd
		return
	default:
		websocket.JSON.Receive(c.con, &test)
	}

}
