package rpn

import "testing"

func TestRpn(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	r := evalRPN(tokens)
	if r != 9 {
		t.Error("rpn calculate error")
	}
}
