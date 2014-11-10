package main

import (
	"testing"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		a   Vec
		b   Vec
		out bool
	}{
		{a: Vec{X: 1.0, Y: 1.0}, b: Vec{X: 1.0, Y: 1.0}, out: true},
		{a: Vec{X: 10.0000001, Y: 1.0}, b: Vec{X: 1.0, Y: 10.0}, out: false},
		{a: Vec{X: 1.0 + precision*0.1, Y: 1.0}, b: Vec{X: 1.0, Y: 1.0}, out: true},
	}

	for _, test := range tests {
		if test.a.Equals(&test.b) != test.out {
			t.Fail()
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a Vec
		b Vec
		r Vec
	}{
		{a: Vec{X: 1.0, Y: 1.0}, b: Vec{X: 1.0, Y: 1.0}, r: Vec{X: 2.0, Y: 2.0}},
		{a: Vec{X: 10.0000001, Y: 1.0}, b: Vec{X: 1.0, Y: 10.0}, r: Vec{X: 11.0000001, Y: 11.0}},
	}

	for _, test := range tests {
		if !test.a.Add(&test.b).Equals(&test.r) {
			t.Errorf("%v + %v = %v Expected %v", test.a, test.b, test.a.Add(&test.b), test.r)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a Vec
		b Vec
		r Vec
	}{
		{a: Vec{X: 1.0, Y: 1.0}, b: Vec{X: 1.0, Y: 1.0}, r: Vec{X: 0.0, Y: 0.0}},
		{a: Vec{X: 10.0000001, Y: 1.0}, b: Vec{X: 1.0, Y: 1.0}, r: Vec{X: 9.0000001, Y: 0.0}},
	}

	for _, test := range tests {
		if !test.a.Sub(&test.b).Equals(&test.r) {
			t.Errorf("%v - %v = %v Expected %v", test.a, test.b, test.a.Sub(&test.b), test.r)
		}
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		a Vec
		f float64
		r Vec
	}{
		{a: Vec{X: 0.0, Y: 0.0}, f: 42.0, r: Vec{X: 0.0, Y: 0.0}},
		{a: Vec{X: 1.0, Y: 0.0}, f: 42.0, r: Vec{X: 42.0, Y: 0.0}},
		{a: Vec{X: 0.0, Y: 1.0}, f: 42.0, r: Vec{X: 0.0, Y: 42.0}},
		{a: Vec{X: 42.1, Y: 42.1}, f: 2.5, r: Vec{X: 105.25, Y: 105.25}},
	}

	for _, test := range tests {
		if !test.a.Scale(test.f).Equals(&test.r) {
			t.Errorf("%v * %v = %v Expected %v", &test.a, test.f, &test.a, &test.r)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		a Vec
		r float64
	}{
		{a: Vec{X: 0.0, Y: 0.0}, r: 0.0},
		{a: Vec{X: 1.0, Y: 0.0}, r: 1.0},
		{a: Vec{X: 0.0, Y: 1.0}, r: 1.0},
		{a: Vec{X: -1.0, Y: 0.0}, r: 1.0},
		{a: Vec{X: 0.0, Y: -1.0}, r: 1.0},
	}

	for _, test := range tests {
		if test.a.length() != test.r {
			t.Errorf("length(%v) = %v Expected %v", &test.a, test.a.length(), test.r)
		}
	}
}

func TestRotate(t *testing.T) {
	//TODO
}
