package decodestring

import "testing"

func TestDecodestring(t *testing.T) {
	s := "3[a2[c]]"
	r := decodeString(s)
	if r != "accaccacc" {
		t.Error("decode string error")
	}
}
