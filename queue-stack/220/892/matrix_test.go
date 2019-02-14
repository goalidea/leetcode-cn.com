package matrix

import "testing"

func TestMatrix(t *testing.T) {

	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	r := updateMatrix(matrix)
	expectation := matrix

	for i := range r {
		for j := range r[i] {
			if expectation[i][j] != r[i][j] {
				t.Error("update matrix error")
			}
		}
	}

}
