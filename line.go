package main

import "math"

type Line [2]Vec

//Slope returns line parameters in slope form
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

//Intersect2 returns the intersection point of Line a and b as Vec
func Intersect2(a Line, b Line) Vec {
	ka, _, pay := a.Slope()
	kb, _, pby := b.Slope()
	x := (pby - pay) / (ka - kb)
	y := kb*x + pby
	return Vec{x, y}
}

func LineAngle(a Line, b Line) float64 {
	wa := math.Atan2(a[0][1]-a[1][1],a[0][0]-a[1][0])
	wb := math.Atan2(b[0][1]-b[1][1],b[0][0]-b[1][0])
	return 2*math.Pi-math.Abs(wa-wb)
}
