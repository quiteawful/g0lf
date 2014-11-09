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

	k = (pm.Y - pn.Y) / (pm.X - pn.X)
	px = pm.X
	py = pm.Y

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
	wa := math.Atan2(a.p1.Y-a.p2.Y, a.p1.X-a.p2.X)
	wb := math.Atan2(b.p1.Y-b.p2.Y, b.p1.X-b.p2.X)
	return 2*math.Pi - math.Abs(wa-wb)
}
