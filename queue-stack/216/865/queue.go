package queue

type MyCircularQueue struct {
	head *node
	tail *node
	size int
}

type node struct {
	next  *node
	value interface{}
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size: k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == 0 {
		return false
	} else {
		n := &node{
			value: value,
		}
		if this.tail == nil {
			this.head = n
			this.tail = n
			this.size--
		} else {
			this.tail.next = n
			this.tail = n
			this.size--
		}
		return true
	}
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == nil {
		return false
	} else {
		if this.head == this.tail {
			this.head, this.tail = nil, nil
			this.size++
		} else {
			this.head = this.head.next
			this.size++
		}
		return true
	}
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head == nil {
		return -1
	}
	if v, ok := this.head.value.(int); ok {
		return v
	} else {
		return -1
	}
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.tail == nil {
		return -1
	}
	if v, ok := this.tail.value.(int); ok {
		return v
	} else {
		return -1
	}
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if (this.head == nil || this.tail == nil) && this.size != 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.head != nil && this.tail != nil && this.size == 0 {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
