package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Bahn struct {
	dim  Vec
	ball *Ball
	hole Vec
}

type Ball struct {
	pos Vec
	vel Vec
}

func main() {

	//websocket
	ss := NewSocketServer("/ws/")
	ss.listen()

	//static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Fatal(http.ListenAndServe(":8181", nil))

	testBahn := new(Bahn)
	testBahn.ball = new(Ball)
	testBahn.ball.pos = Vec{5.0, 15.0}
	testBahn.dim = Vec{80.0, 30.0}
	testBahn.hole = Vec{75.0, 15.0}
	//testBahn.Print()

	testline := Line{Vec{0, 0}, Vec{3, 1}}
	wall := Line{Vec{0, 10}, Vec{10, 10}}
	fmt.Println(Intersect2(wall, testline))
	fmt.Println(Intersect2(testline, wall))

	frameNS := time.Duration(int(1e9) / 30)
	clk := time.NewTicker(frameNS)

	for {
		select {
		case <-clk.C:
			//stuff
		}
	}
}

func (b *Bahn) Print() {
	var i, j float64
	for i = 0; i < b.dim.x; i++ {
		fmt.Print("_")
	}
	fmt.Print("\n")

	for i = 0; i < b.dim.y; i++ {
		for j = 0; j < b.dim.x; j++ {
			if i == b.ball.pos.y && j == b.ball.pos.x {
				fmt.Print("x")
				continue
			}
			if i == b.hole.y && j == b.hole.x {
				fmt.Print("o")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Print("|", i, "\n")
	}

	for i = 0; i < b.dim.x; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	fmt.Println(b.dim.x, b.dim.y)
}
