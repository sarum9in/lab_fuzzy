package main

import (
	"math"
)

type GaussFunction struct {
	C float64
	S float64
}

func (f *GaussFunction) Value(x float64) float64 {
	inner := (x - f.C) / f.S
	return math.Exp(-(inner * inner))
}

type TriangleFunction struct {
	C float64
	B float64
}

func (f *TriangleFunction) Value(x float64) float64 {
	if x <= f.C {
		return (x-f.C)/f.B + 1
	} else {
		return (f.C-x)/f.B + 1
	}
}

type Env struct {
	Fx1 GaussFunction
	Fx2 GaussFunction
	Fy  TriangleFunction
}

func main() {
	envs := []Env{
		Env{
			Fx1: GaussFunction{0, 3},
			Fx2: GaussFunction{0, 4.01},
			Fy:  TriangleFunction{0, 3.62},
		},
		Env{
			Fx1: GaussFunction{3.64, 4.86},
			Fx2: GaussFunction{3.16, 3.35},
			Fy:  TriangleFunction{4.95, 3.66},
		},
		Env{
			Fx1: GaussFunction{8, 7.59},
			Fx2: GaussFunction{8, 4.64},
			Fy:  TriangleFunction{8, 1.43},
		},
		Env{
			Fx1: GaussFunction{0, 6.49},
			Fx2: GaussFunction{8, 4.65},
			Fy:  TriangleFunction{4, 2.78},
		},
		Env{
			Fx1: GaussFunction{8, 7.13},
			Fx2: GaussFunction{0, 4.40},
			Fy:  TriangleFunction{4.55, 3.31},
		},
	}
}
