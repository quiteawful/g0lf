package main

import "log"

type Game struct {
	WebSocket *SocketServer
	Opts      Options
	CurrSt    State
	Tr        *Track
	msg       chan GMessage
}

type Options struct {
	MaxPlayers int
}

type GMessage struct {
	MsgType string `json:"type"`
	X       float64
	Y       float64
}
type State struct {
	PlayerCount int
}

func (g *Game) NewGame(name string, maxplayers int) {
	g.WebSocket = NewSocketServer(name, g)
	g.Opts.MaxPlayers = maxplayers
	g.msg = make(chan GMessage)
	go g.WebSocket.listen()
}

func (g *Game) SendAllPlayers(i interface{}) {
	for _, p := range g.WebSocket.clients {
		p.JSON <- i
	}
}

func (g *Game) StartGame() {
	log.Println("Entering game loop")
	for {
		select {
		case <-newConn:
			g.CurrSt.PlayerCount++
			if g.CurrSt.PlayerCount == g.Opts.MaxPlayers {
				// Put level in a map to have a way to identify the message type
				g.SendAllPlayers(map[string]interface{}{"leveldata": g.Tr})
			}

		case msg := <-g.msg:
			log.Println("Game got JSON:", msg.MsgType)
		case derp := <-dummy:
			_ = derp
			panic("should not be reached")

		}
	}
}
