package main

import "math"

type Line struct {
	p1 Vec
	p2 Vec
}

//Slope returns line parameters in slope form
func (l *Line) Slope() (k float64, px float64, py float64) {

	pm := Vec{0, 0}
	pn := Vec{0, 0}

	if l.p1.length() > l.p2.length() {
		pm, pn = l.p2, l.p1
	} else {
		pm, pn = l.p1, l.p2
	}

	k = (pm.y - pn.y) / (pm.x - pn.x)
	px = pm.x
	py = pm.y

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
	wa := math.Atan2(a.p1.y-a.p2.y, a.p1.x-a.p2.x)
	wb := math.Atan2(b.p1.y-b.p2.y, b.p1.x-b.p2.x)
	return 2*math.Pi - math.Abs(wa-wb)
}
