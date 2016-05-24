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

type Rule struct {
	Conditions []Condition
	Conclusion Conclusion
}

type Condition struct {
	InputID int
	Term    string
}

type Conclusion struct {
	Term string
}

type Env struct {
	Fx    []GaussFunction
	Fy    TriangleFunction
	Rules []Rule
}

func main() {
	envs := []Env{
		Env{
			Fx: []GaussFunction{
				GaussFunction{0, 3},
				GaussFunction{0, 4.01},
			},
			Fy: TriangleFunction{0, 3.62},
		},
		Env{
			Fx: []GaussFunction{
				GaussFunction{3.64, 4.86},
				GaussFunction{3.16, 3.35},
			},
			Fy: TriangleFunction{4.95, 3.66},
		},
		Env{
			Fx: []GaussFunction{
				GaussFunction{8, 7.59},
				GaussFunction{8, 4.64},
			},
			Fy: TriangleFunction{8, 1.43},
		},
		Env{
			Fx: []GaussFunction{
				GaussFunction{0, 6.49},
				GaussFunction{8, 4.65},
			},
			Fy: TriangleFunction{4, 2.78},
		},
		Env{
			Fx: []GaussFunction{
				GaussFunction{8, 7.13},
				GaussFunction{0, 4.40},
			},
			Fy: TriangleFunction{4.55, 3.31},
		},
	}
	for i, _ := range envs {
		envs[i].Rules = []Rule{
			Rule{
				[]Condition{
					Condition{0, "A11"},
					Condition{1, "A12"},
				},
				Conclusion{"B1"},
			},
			Rule{
				[]Condition{
					Condition{0, "A21"},
					Condition{1, "A22"},
				},
				Conclusion{"B2"},
			},
			Rule{
				[]Condition{
					Condition{0, "A31"},
					Condition{1, "A32"},
				},
				Conclusion{"B3"},
			},
			Rule{
				[]Condition{
					Condition{0, "A11"},
					Condition{1, "A32"},
				},
				Conclusion{"B2"},
			},
			Rule{
				[]Condition{
					Condition{0, "A31"},
					Condition{1, "A12"},
				},
				Conclusion{"B2"},
			},
		}
	}
	x := []float64{4.90, 2.44}
	for _, env := range envs {
		// fuzzification
		fuzzified := make([]float64, 0, 10)
		for _, rule := range env.Rules {
			for _, condition := range rule.Conditions {
				id := condition.InputID
				fuzzified = append(fuzzified, env.Fx[id].Value(x[id]))
			}
		}
		// aggregation
		truth := make([]float64, len(env.Rules))
		i := 0
		for ruleID, rule := range env.Rules {
			truth[ruleID] = 1
			for _, condition := range rule.Conditions {
				truth[ruleID] = math.Min(truth[ruleID], fuzzified[i])
				i++
			}
		}
	}
}
