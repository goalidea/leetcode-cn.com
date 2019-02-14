package targetsum

import "testing"

func TestTargetsum(t *testing.T) {
	nums, S := []int{1, 1, 1, 1, 1}, 3
	r := findTargetSumWays(nums, S)
	if r != 5 {
		t.Error("func findtargetsumways count error")
	}
}
