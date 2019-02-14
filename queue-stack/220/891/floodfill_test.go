package floodfill

import "testing"

func TestFloodFill(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, newColor := 1, 1, 2
	expection := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

	r := floodFill(image, sr, sc, newColor)
	for ri := range r {
		for rj := range r[ri] {
			if r[ri][rj] != expection[ri][rj] {
				t.Error("func floodFill error")
			}
		}
	}
}
