package flatten

type ListNode struct {
	Val   int
	Prev  *ListNode
	Next  *ListNode
	Child *ListNode
}

func flatten(head *ListNode) *ListNode {
	flattentail(head)
	return head
}

func flattentail(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Child == nil {
		if head.Next == nil {
			return head
		}
		return flattentail(head.Next)
	} else {
		child, next := head.Child, head.Next
		head.Next, child.Prev, head.Child = child, head, nil
		headtail := flattentail(child)
		if next != nil {
			headtail.Next = next
			next.Prev = headtail
			return flattentail(next)
		}
		return headtail
	}
}
