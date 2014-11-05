package main

import "math"

type Vec struct {
	x float64
	y float64
}

const precision = 0.0000001

// Add adds vector a to v
func (v *Vec) Add(a *Vec) *Vec {
	v.x = v.x + a.x
	v.y = v.y + a.y
	return v
}

// Sub substracts vector a from v
func (v *Vec) Sub(a *Vec) *Vec {
	v.x = v.x - a.x
	v.y = v.y - a.y
	return v
}

// Equals compares two vectors and returns true if the difference
// is smaller than the defined precision
func (v *Vec) Equals(a *Vec) bool {
	diff := v.Sub(a)
	if diff.x < precision && diff.y < precision {
		return true
	}
	return false
}

// Dotp calculates the dot product of two vectors (Skalarprodukt)
func Dotp(a, b *Vec) float64 {
	return a.x*b.x + a.y*b.y
}

// length returns the length of a vector
func (v *Vec) length() float64 {
	return math.Sqrt(Dotp(v, v))
}

// Rotate vector by rad radians
func (v *Vec) Rotate(rad float64) {
	nx := v.x*math.Cos(rad) - v.y*math.Sin(rad)
	ny := v.x*math.Sin(rad) + v.y*math.Cos(rad)
	v.x, v.y = nx, ny
}
