package V1

import (
	"math"
)

func gaussPartial(a0 [][]float64, b0 []float64) ([]float64, error) {
	m := len(b0)
	a := make([][]float64, m)
	for i, ai := range a0 {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	for k := range a {
		iMax := 0
		max := -1.
		for i := k; i < m; i++ {
			row := a[i]
			s := -1.
			for j := k; j < m; j++ {
				x := math.Abs(row[j])
				if x > s {
					s = x
				}
			}
			if abs := math.Abs(row[k]) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		if a[iMax][k] == 0 {
			return nil, newError(errorTypSolving, errorCriticalLevelPartial, "Gaussian Singular")
		}
		a[k], a[iMax] = a[iMax], a[k]
		for i := k + 1; i < m; i++ {
			for j := k + 1; j <= m; j++ {
				a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
			}
			a[i][k] = 0
		}
	}
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x, nil
}
