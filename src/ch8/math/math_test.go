package math

import "testing"

type testPairAverage struct {
	values  []float64
	average float64
}

type testPairMax struct {
	values []float64
	max    float64
}

type testPairMin struct {
	values []float64
	max    float64
}

var averageTests = []testPairAverage{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

var maxTests = []testPairMax{
	{[]float64{1, 2}, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
	{[]float64{1, -3, 0, 4, 2, -2}, 4},
	{[]float64{-1, -3, -2.4, -4, -2, -2}, -1},
	{[]float64{}, 0},
}

var minTests = []testPairMax{
	{[]float64{1, 2}, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, -1},
	{[]float64{1, -3, 0, 4, 2, -2}, -3},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)

		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)

		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
