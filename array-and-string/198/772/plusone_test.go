package plusone

import "testing"

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 3}

	r := plusOne(digits)

	expectation := []int{1, 2, 4}

	for i := range r {
		if r[i] != expectation[i] {
			t.Error("plus one error")
		}
	}
}
