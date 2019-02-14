package temperature

import "testing"

func TestTemperature(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := []int{1, 1, 4, 2, 1, 1, 0, 0}
	r := dailyTemperatures(T)
	for i := 0; i < 8; i++ {
		if r[i] != result[i] {
			t.Error("func dailytemperatures return error")
		}
	}
}
