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
	envs := []Env{
		Env{
			Rx1: GaussRule{0, 3},
			Rx2: GaussRule{0, 4.01},
			Ry:  TriangleRule{0, 3.62},
		},
		Env{
			Rx1: GaussRule{3.64, 4.86},
			Rx2: GaussRule{3.16, 3.35},
			Ry:  TriangleRule{4.95, 3.66},
		},
		Env{
			Rx1: GaussRule{8, 7.59},
			Rx2: GaussRule{8, 4.64},
			Ry:  TriangleRule{8, 1.43},
		},
		Env{
			Rx1: GaussRule{0, 6.49},
			Rx2: GaussRule{8, 4.65},
			Ry:  TriangleRule{4, 2.78},
		},
		Env{
			Rx1: GaussRule{8, 7.13},
			Rx2: GaussRule{0, 4.40},
			Ry:  TriangleRule{4.55, 3.31},
		},
	}
}
