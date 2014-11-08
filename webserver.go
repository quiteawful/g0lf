package main

import (
	"log"

	"code.google.com/p/go.net/websocket"
)

func wsHandler(ws *websocket.Conn) {
	log.Println("new connection")
}
