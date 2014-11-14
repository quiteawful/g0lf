package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bahn struct {
	Dim  Vec
	Ball *Ball
	Hole Vec
}

type Ball struct {
	Pos Vec
	Vel Vec
}

var newConn = make(chan *Client)
var dummy = make(chan string)

func main() {

	//websocket
	ss := NewSocketServer("/ws/")
	go ss.listen()

	//static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	testBahn := new(Bahn)
	testBahn.Ball = new(Ball)
	testBahn.Ball.Pos = Vec{5.0, 15.0}
	testBahn.Dim = Vec{80.0, 30.0}
	testBahn.Hole = Vec{75.0, 15.0}
	b, err := json.MarshalIndent(testBahn, "", "	")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(b))
	//testBahn.Print()

	testline := Line{Vec{0, 0}, Vec{3, 1}}
	wall := Line{Vec{0, 10}, Vec{10, 10}}
	fmt.Println(Intersect2(wall, testline))
	fmt.Println(Intersect2(testline, wall))
	go func() {
		log.Fatal(http.ListenAndServe(":8181", nil))
	}()
	log.Println("Entering main loop")
	for {
		select {
		case cl := <-newConn:
			//cl.msg <- "Hallo neuer client!"
			// Put level in a map to have a way to identify the message type
			cl.JSON <- map[string]interface{}{"leveldata": testBahn}
		case derp := <-dummy:
			_ = derp
			panic("should not be reached")

		}

	}
}

func (b *Bahn) Print() {
	var i, j float64
	for i = 0; i < b.Dim.X; i++ {
		fmt.Print("_")
	}
	fmt.Print("\n")

	for i = 0; i < b.Dim.Y; i++ {
		for j = 0; j < b.Dim.X; j++ {
			if i == b.Ball.Pos.Y && j == b.Ball.Pos.X {
				fmt.Print("x")
				continue
			}
			if i == b.Hole.Y && j == b.Hole.X {
				fmt.Print("o")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Print("|", i, "\n")
	}

	for i = 0; i < b.Dim.X; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	fmt.Println(b.Dim.X, b.Dim.Y)
}
