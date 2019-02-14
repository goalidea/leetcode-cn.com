package designsinglylist

type MyLinkedList struct {
	val  interface{}
	len  int
	head *MyLinkedList
	tail *MyLinkedList
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len {
		return -1
	}

	tmp := new(MyLinkedList)
	tmp = this.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}

	v, ok := tmp.val.(int)
	if !ok {
		panic("get val error")
	}
	return v
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = this
		this.tail = this
		this.val = val
		this.len++
		return
	}
	this.len++
	this.head = &MyLinkedList{
		val:  val,
		next: this.head,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.tail == nil {
		this.head = this
		this.tail = this
		this.val = val
		this.len++
		return
	}
	this.len++
	tmp := &MyLinkedList{}
	tmp.val = val
	this.tail.next = tmp
	this.tail = tmp
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len || index < 0 {
		return
	}
	if index == this.len {
		this.AddAtTail(val)
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	this.len++
	tmp := new(MyLinkedList)
	tmp = this.head
	for i := 0; i < index-1; i++ {
		tmp = tmp.next
	}
	tmpcur := &MyLinkedList{}
	tmpcur.val = val
	tmpcur.next = tmp.next
	tmp.next = tmpcur
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}

	if this.len == 1 {
		this.val = nil
		this.len--
		this.head = nil
		this.tail = nil
		this.next = nil
		return
	}

	// find previous node which node need to delete
	tmp := new(MyLinkedList)
	tmp = this.head
	for i := 0; i < index-1; i++ {
		tmp = tmp.next
	}

	if index == this.len-1 {
		this.len--
		this.tail = tmp
		tmp.next = nil
		return
	}

	this.len--
	tmp.next = tmp.next.next
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
