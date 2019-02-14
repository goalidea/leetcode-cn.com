package queueTstack

import "testing"

func TestStackToQueue(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	if r := stack.Top(); r != 2 {
		t.Error("get top of stack error")
	}
	if r := stack.Pop(); r != 2 {
		t.Error("remove top of stack error")
	}
	if stack.Empty() {
		t.Error("stack has one element but empty error")
	}
}
