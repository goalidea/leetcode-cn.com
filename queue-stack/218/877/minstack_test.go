package minstack

import "testing"

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	if mini1 := ms.GetMin(); mini1 != -3 {
		t.Error("get mini value of stack error")
	}
	ms.Pop()
	if tp := ms.Top(); tp != 0 {
		t.Error("get top value of stack error")
	}
	if mini2 := ms.GetMin(); mini2 != -2 {
		t.Error("get mini value of stack error")
	}
}
