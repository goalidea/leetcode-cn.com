package pivotindex

import "testing"

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}

	expectation := 3

	r := pivotIndex(nums)

	if r != expectation {
		t.Error("find pivot index error")
	}
}
