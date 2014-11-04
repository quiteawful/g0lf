package main

import "math"

type Vec [2]float64

const precision = 0.0000001

// Add adds vector a to v
func (v *Vec) Add(a *Vec) *Vec {
	v[0] = v[0] + a[0]
	v[1] = v[1] + a[1]
	return v
}

// Sub substracts vector a from v
func (v *Vec) Sub(a *Vec) *Vec {
	v[0] = v[0] - a[0]
	v[1] = v[1] - a[1]
	return v
}

// Equals compares two vectors and returns true if the difference
// is smaller than the defined precision
func (v *Vec) Equals(a *Vec) bool {
	diff := v.Sub(a)
	if diff[0] < precision && diff[1] < precision {
		return true
	}
	return false
}

// Dotp calculates the dot product of two vectors (Skalarprodukt)
func Dotp(a, b *Vec) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

// length returns the length of a vector
func (v *Vec) length() float64 {
	return math.Sqrt(Dotp(v, v))
}

func (v *Vec) rotate(rad float64){
	nx := v[0] * math.Cos(rad) - v[1] * math.Sin(rad)
	ny := v[0] * math.Sin(rad) + v[1] * math.Cos(rad)
	v[0], v[1] = nx, ny
}
