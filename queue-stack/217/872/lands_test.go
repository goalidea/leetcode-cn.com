package lands

import "testing"

func TestNumIslands(t *testing.T) {
	g := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	r := numIslands(g)
	if r != 1 {
		t.Error("num of lands calculate error")
	}
}
