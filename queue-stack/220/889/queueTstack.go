package queueTstack

import "container/list"

type MyStack struct {
	queue0 *list.List
	queue1 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue0: list.New(),
		queue1: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.queue0.Len() == 0 && this.queue1.Len() == 0 {
		this.queue0.PushBack(x)
		return
	}
	if this.queue0.Len() > 0 {
		this.queue0.PushBack(x)
		return
	}
	if this.queue1.Len() > 0 {
		this.queue1.PushBack(x)
		return
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.queue0.Len() > 0 {
		queueTbrothers(this.queue0, this.queue1)
		return this.queue0.Remove(this.queue0.Front()).(int)
	}
	queueTbrothers(this.queue1, this.queue0)
	return this.queue1.Remove(this.queue1.Front()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	tmp := this.Pop()
	this.Push(tmp)
	return tmp
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue0.Len() == 0 && this.queue1.Len() == 0
}

func queueTbrothers(start *list.List, end *list.List) {
	for start.Len() > 0 {
		if start.Len() == 1 {
			break
		}
		end.PushBack(start.Remove(start.Front()))
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
