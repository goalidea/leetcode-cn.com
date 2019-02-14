package queue

import "testing"

func TestQueue(t *testing.T) {
	//queue{nil, nil, 3}
	queue := Constructor(3)

	//queue{&{nil, 1}, &{nil, 1}, 2}
	if ok := queue.EnQueue(1); !ok {
		t.Error("use value 1 enqueue error")
	}

	//queue{&{&{nil, 2}, 1}, &{nil, 2}, 1}
	if ok := queue.EnQueue(2); !ok {
		t.Error("use value 2 enqueue error")
	}

	//queue{&{&{&{nil, 3}, 2}, 1}, &{&{nil, 3}, 2}, &{nil, 3}, 0}
	if ok := queue.EnQueue(3); !ok {
		t.Error("use value 3 enqueue error")
	}

	//queue{&{&{&{nil, 3}, 2}, 1}, &{&{nil, 3}, 2}, &{nil, 3}, 0}
	if ok := queue.EnQueue(4); ok {
		t.Error("queue isfull but enqueue success")
	}

	//queue{&{&{&{nil, 3}, 2}, 1}, &{&{nil, 3}, 2}, &{nil, 3}, 0}
	if queue.Rear() != 3 {
		t.Error("get back of queue error")
	}

	//queue{&{&{&{nil, 3}, 2}, 1}, &{&{nil, 3}, 2}, &{nil, 3}, 0}
	if !queue.IsFull() {
		t.Error("now queue isfull but not")
	}

	//queue{&{&{nil, 3}, 2}, &{nil, 3}, 1}
	if !queue.DeQueue() {
		t.Error("dequeue error")
	}

	//queue{&{&{&{nil, 4}, 3}, 2}, &{&{ni,4}, 3}, &{nil, 4}, 0}
	if ok := queue.EnQueue(4); !ok {
		t.Error("use value 4 enqueue error")
	}

	//queue{&{&{&{nil, 4}, 3}, 2}, &{&{ni,4}, 3}, &{nil, 4}, 0}
	if queue.Rear() != 4 {
		t.Error("get back of queue error")
	}
}
