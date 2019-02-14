package lock

import (
	"testing"
)

func TestLock(t *testing.T) {
	dead, targ := []string{"0201", "0101", "0102", "1212", "2002"}, "0202"
	r := openLock(dead, targ)
	if r != 6 {
		t.Error("openlock fail")
	}
}
