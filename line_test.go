package main

import (
	"testing"
)

func TestSlope(t *testing.T) {

}

func TestIntersect2(t *testing.T) {

}

func TestLineAngle(t *testing.T) {

}

func TestParallel(t *testing.T) {

	tests := []struct {
		a   Line
		b   Line
		out bool
	}{
		{a: Line{p1: Vec{1, 1}, p2: Vec{2, 2}}, b: Line{p1: Vec{3, 3}, p2: Vec{4, 4}}, out: true},
		{a: Line{p1: Vec{1, 1}, p2: Vec{2, 2}}, b: Line{p1: Vec{3, 3}, p2: Vec{4, 5}}, out: false},
	}

	for _, test := range tests {
		if Parallel(test.a, test.b) != test.out {
			t.Errorf("Parallel(%v, %v) = %v, Expected %v", test.a, test.b, Parallel(test.a, test.b), test.out)
		}
	}
}
