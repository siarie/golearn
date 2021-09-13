package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var test = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 2, 3}, 2},
	{[]float64{1, 2, 2, 1, 1}, 1.4},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range test {
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
