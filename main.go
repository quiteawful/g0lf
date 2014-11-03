package main

import "fmt"

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
	fmt.Println("asdf")

	testBahn := new(Bahn)
	testBahn.ball = new(Ball)
	testBahn.ball.pos = Vec{5.0, 15.0}
	testBahn.dim = Vec{80.0, 30.0}
	testBahn.hole = Vec{75.0, 15.0}
	testBahn.Print()

}

func (b *Bahn) Print() {
	var i, j float64
	for i = 0; i < b.dim[0]; i++ {
		fmt.Print("_")
	}
	fmt.Print("\n")

	for i = 0; i < b.dim[1]; i++ {
		for j = 0; j < b.dim[0]; j++ {
			if i == b.ball.pos[1] && j == b.ball.pos[0] {
				fmt.Print("x")
				continue
			}
			if i == b.hole[1] && j == b.hole[0] {
				fmt.Print("o")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Print("|", i, "\n")
	}

	for i = 0; i < b.dim[0]; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	fmt.Println(b.dim[0], b.dim[1])
}
