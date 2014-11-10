package main

import "math"

type Vec struct {
	X float64
	Y float64
}

const precision = 0.0000001

// Add adds vector a to v
func (v *Vec) Add(a *Vec) *Vec {
	v.X = v.X + a.X
	v.Y = v.Y + a.Y
	return v
}

// Sub substracts vector a from v
func (v *Vec) Sub(a *Vec) *Vec {
	v.X = v.X - a.X
	v.Y = v.Y - a.Y
	return v
}

// Equals compares two vectors and returns true if the difference
// is smaller than the defined precision
func (v *Vec) Equals(a *Vec) bool {
	diff := v.Sub(a)
	if diff.X < precision && diff.Y < precision {
		return true
	}
	return false
}

// Dotp calculates the dot product of two vectors (Skalarprodukt)
func Dotp(a, b *Vec) float64 {
	return a.X*b.X + a.Y*b.Y
}

// length returns the length of a vector
func (v *Vec) length() float64 {
	return math.Sqrt(Dotp(v, v))
}

// Rotate vector by rad radians
func (v *Vec) Rotate(rad float64) {
	nx := v.X*math.Cos(rad) - v.Y*math.Sin(rad)
	ny := v.X*math.Sin(rad) + v.Y*math.Cos(rad)
	v.X, v.Y = nx, ny
}

// Scale scales a vector v by factor f
func (v *Vec) Scale(f float64) *Vec {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}
