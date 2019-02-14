package square

import (
	"testing"
)

func TestSquare(t *testing.T) {
	n := 12
	r := numSquares(n)
	if r != 3 {
		t.Error("count error")
	}
}
