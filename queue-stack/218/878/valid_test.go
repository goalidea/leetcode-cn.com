package valid

import "testing"

func TestValid(t *testing.T) {
	s := "()"
	r := isValid(s)
	if !r {
		t.Error("func isvalid error")
	}
}
