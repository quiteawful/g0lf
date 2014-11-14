package main

import "log"

type Game struct {
	WebSocket *SocketServer
	Opts      Options
	CurrSt    State
	Tr        *Track
}

type Options struct {
	MaxPlayers int
}

type State struct {
	PlayerCount int
}

func (g *Game) NewGame(name string, maxplayers int) {
	g.WebSocket = NewSocketServer(name)
	g.Opts.MaxPlayers = maxplayers
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
		case cl := <-newConn:
			g.CurrSt.PlayerCount++
			if g.CurrSt.PlayerCount == g.Opts.MaxPlayers {
				// Put level in a map to have a way to identify the message type
				cl.JSON <- map[string]interface{}{"leveldata": g.Tr}
				g.SendAllPlayers(map[string]interface{}{"leveldata": g.Tr})
			}
		case derp := <-dummy:
			_ = derp
			panic("should not be reached")

		}
	}
}
