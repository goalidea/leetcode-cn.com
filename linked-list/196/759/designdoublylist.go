package designdoublylist

type MyLinkedList struct {
	size int
	head *node
	tail *node
}

type node struct {
	val  int
	prev *node
	next *node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size-1 {
		return -1
	}

	tmp := this.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}

	return tmp.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &node{
		val: val,
	}

	if this.size == 0 {
		this.size++
		this.head = tmp
		this.tail = tmp
		return
	}

	this.size++
	tmp.next = this.head
	this.head.prev = tmp
	this.head = tmp
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := &node{
		val: val,
	}

	if this.size == 0 {
		this.size++
		this.head = tmp
		this.tail = tmp
		return
	}

	this.size++
	tmp.prev = this.tail
	this.tail.next = tmp
	this.tail = tmp
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	this.size++

	tmp := this.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}

	tmpnode := &node{
		val: val,
	}

	tmpnode.next = tmp
	tmpnode.prev = tmp.prev
	tmp.prev.next = tmpnode
	tmp.prev = tmpnode
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size-1 || this.size == 0 {
		return
	}

	if index == 0 {
		this.size--
		this.head.next.prev = this.head.prev
		this.head = this.head.next
		return
	}

	if index == this.size-1 {
		this.size--
		this.tail.prev.next = this.tail.next
		this.tail = this.tail.prev
		return
	}

	tmp := this.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}

	tmp.prev.next = tmp.next
	tmp.next.prev = tmp.prev
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
