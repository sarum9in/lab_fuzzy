package main

import (
	"math"
)

type GaussRule struct {
	C float64
	S float64
}

func (r *GaussRule) Value(x float64) float64 {
	inner := (x - r.C) / r.S
	return math.Exp(-(inner * inner))
}

type TriangleRule struct {
	C float64
	B float64
}

func (r *TriangleRule) Value(x float64) float64 {
	if x <= r.C {
		return (x-r.C)/r.B + 1
	} else {
		return (r.C-x)/r.B + 1
	}
}

type Env struct {
	Rx1 GaussRule
	Rx2 GaussRule
	Ry  TriangleRule
}

func main() {
}
