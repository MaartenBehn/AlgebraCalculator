package AlgebraCalculator

import "testing"

type testCase struct {
	a [][]float64
	b []float64
}

var tc = testCase{
	a: [][]float64{
		{1.00, 0.00, 0.00, 0.00, 0.00, 0.00},
		{1.00, 0.63, 0.39, 0.25, 0.16, 0.10},
		{1.00, 1.26, 1.58, 1.98, 2.49, 3.13},
		{1.00, 1.88, 3.55, 6.70, 12.62, 23.80},
		{1.00, 2.51, 6.32, 15.88, 39.90, 100.28},
		{1.00, 3.14, 9.87, 31.01, 97.41, 306.02},
	},
	b: []float64{-0.01, 0.61, 0.91, 0.99, 0.60, 0.02},
}

var tc2 = testCase{
	a: [][]float64{
		{6, 1, 2},
		{5, 3, 3},
		{3, 2, 1},
	},
	b: []float64{1, 4, 14},
}

// result from above test case turns out to be correct to this tolerance.
const tolerance = 1e-14

func TestGaussianElimination(t *testing.T) {
	x, err := gaussPartial(tc2.a, tc2.b)
	if err != nil {
		t.Error(err)
	}
	t.Log(x)

	x, err = gaussPartial(tc.a, tc.b)
	if err != nil {
		t.Error(err)
	}
	t.Log(x)
}
