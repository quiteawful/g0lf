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
		{a: Vec{x: 1.0, y: 1.0}, b: Vec{x: 1.0, y: 1.0}, out: true},
		{a: Vec{x: 10.0000001, y: 1.0}, b: Vec{x: 1.0, y: 10.0}, out: false},
		// more testcases
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
		{a: Vec{x: 1.0, y: 1.0}, b: Vec{x: 1.0, y: 1.0}, r: Vec{x: 2.0, y: 2.0}},
		{a: Vec{x: 10.0000001, y: 1.0}, b: Vec{x: 1.0, y: 10.0}, r: Vec{x: 11.0000001, y: 11.0}},
		// more testcases
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
		{a: Vec{x: 1.0, y: 1.0}, b: Vec{x: 1.0, y: 1.0}, r: Vec{x: 0.0, y: 0.0}},
		{a: Vec{x: 10.0000001, y: 1.0}, b: Vec{x: 1.0, y: 1.0}, r: Vec{x: 9.0000001, y: 0.0}},
		// more testcases
	}

	for _, test := range tests {
		if !test.a.Sub(&test.b).Equals(&test.r) {
			t.Errorf("%v - %v = %v Expected %v", test.a, test.b, test.a.Sub(&test.b), test.r)
		}
	}
}
