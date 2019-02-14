package designsinglylist

import "testing"

func TestDesignList(t *testing.T) {
	lt := Constructor()
	lt.AddAtHead(1)
	lt.AddAtTail(3)
	lt.AddAtIndex(1, 2)
	r1 := lt.Get(1)
	lt.DeleteAtIndex(1)
	r2 := lt.Get(1)

	if r1 != 2 && r2 != 3 {
		t.Error("design linked list error")
	}
}
