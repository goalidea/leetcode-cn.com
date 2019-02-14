package designdoublylist

import "testing"

func TestDesignDoublyList(t *testing.T) {
	dlist := Constructor()
	dlist.AddAtHead(1)
	dlist.AddAtTail(3)
	dlist.AddAtIndex(1, 2)
	r1 := dlist.Get(1)
	dlist.DeleteAtIndex(1)
	r2 := dlist.Get(1)

	if r1 != 2 || r2 != 3 {
		t.Error("design doubly linked list error")
	}
}
