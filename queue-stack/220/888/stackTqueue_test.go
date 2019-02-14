package stackTqueue

import "testing"

func TestStackToQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	if r := q.Peek(); r != 1 {
		t.Error("queue get font element error")
	}
	if r := q.Pop(); r != 1 {
		t.Error("queue remove font element error")
	}
	if q.Empty() {
		t.Error("queue has one element but empty error")
	}
}
