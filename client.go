package main

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

var gid int = 0

type Client struct {
	id     int
	con    *websocket.Conn
	server *SocketServer
	quitCh chan bool
	msg    chan string
	JSON   chan interface{}
}

func NewClient(ws *websocket.Conn, s *SocketServer) *Client {

	gid++
	return &Client{
		gid,
		ws,
		s,
		make(chan bool),
		make(chan string),
		make(chan interface{}),
	}
}

func (c *Client) listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) listenWrite() {
	log.Println("Starting Client.listenWrite()")
	for {
		select {
		case msg := <-c.msg:
			if err := websocket.Message.Send(c.con, msg); err != nil {
				log.Println("Error sending message: ", err.Error())
				continue
			}
		case j := <-c.JSON:
			if err := websocket.JSON.Send(c.con, j); err != nil {
				log.Println(err.Error())
				continue
			}
		case <-c.quitCh:
			c.server.Del(c)
			c.con.Close()
			return
		}
	}
}

func (c *Client) Close() {
	c.quitCh <- true
}

func (c *Client) listenRead() {
	log.Println("Starting Client.listenRead()")
	var test string
	for {
		select {
		case <-c.quitCh:
			c.server.Del(c)
			return
		default:
			if err := websocket.Message.Receive(c.con, &test); err != nil {
				if err == io.EOF {
					c.quitCh <- true
				} else {
					log.Println(err.Error())
				}
			} else {
				log.Println("got: ", test)
			}
		}
	}
}
