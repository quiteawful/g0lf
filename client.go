package main

import (
	"io"
	"log"
	"time"

	"code.google.com/p/go.net/websocket"
)

var gid int = 0

type Client struct {
	id     int
	con    *websocket.Conn
	server *SocketServer
	quitCh chan bool
	msg    chan string
}

func NewClient(ws *websocket.Conn, s *SocketServer) *Client {

	gid++
	return &Client{
		gid,
		ws,
		s,
		make(chan bool),
		make(chan string),
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
				return
			}
			log.Println("sent message")
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
	go func() {
		time.Sleep(2 * time.Second)
		c.msg <- "Test123"
	}()
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
