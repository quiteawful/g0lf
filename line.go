package main

import "fmt"

type Line [2]Vec

func (l *Line) Slope() (k float64, px float64, py float64) {

	p1 := Vec{0, 0}
	p2 := Vec{0, 0}

	if l[0].length() > l[1].length() {
		p1, p2 = l[1], l[0]
	} else {
		p1, p2 = l[0], l[1]
	}

	k = (p2[1] - p1[1]) / (p2[0] - p1[0])
	px = p1[0]
	py = p1[1]

	return
}

func Intersect2(a Line, b Line) Vec {
	ka, pax, pay := a.Slope()
	kb, pbx, pby := b.Slope()
	fmt.Println(ka, pax, pay)
	fmt.Println(kb, pbx, pby)
	x := (pby - pay) / (ka - kb)
	y := kb*x + pby
	//	x := (kb*pbx - pby - ka*pax + pay)
	//	y := kb*(x-pbx) + pby

	return Vec{x, y}
}
