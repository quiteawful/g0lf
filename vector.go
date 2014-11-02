package main

type Vec [2]float64

const precision = 0.0000001

func (v *Vec) Add(a *Vec) *Vec {
	v[0] = v[0] + a[0]
	v[1] = v[1] + a[1]
	return v
}

func (v *Vec) Del(a *Vec) *Vec {
	v[0] = v[0] - a[0]
	v[1] = v[1] - a[1]
	return v
}

func (v *Vec) Equals(a *Vec) bool {
	diff := v.Del(a)
	if diff[0] < precision && diff[1] < precision {
		return true
	}
	return false
}
