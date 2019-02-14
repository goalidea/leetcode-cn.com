package stackTqueue

import "container/list"

type MyQueue struct {
	store  *list.List
	search *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		store:  list.New(),
		search: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.store.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.search.Len() == 0 {
		storeTsearch(this.store, this.search)
	}
	return this.search.Remove(this.search.Back()).(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.search.Len() == 0 {
		storeTsearch(this.store, this.search)
	}
	return this.search.Back().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.store.Len() == 0 && this.search.Len() == 0 {
		return true
	}
	return false
}

func storeTsearch(store *list.List, search *list.List) {
	for store.Len() > 0 {
		search.PushBack(store.Remove(store.Back()))
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
