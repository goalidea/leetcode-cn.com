package flatten

import "testing"

func TestFlatten(t *testing.T) {
	head := &ListNode{}
	head.Val = 1
	head.Next = &ListNode{Val: 2, Prev: head}
	head.Next.Next = &ListNode{Val: 3, Prev: head.Next}
	head.Next.Next.Next = &ListNode{Val: 4, Prev: head.Next.Next}
	head.Next.Next.Next.Next = &ListNode{Val: 5, Prev: head.Next.Next.Next}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6, Prev: head.Next.Next.Next.Next}

	headchild1 := &ListNode{Val: 7}
	headchild1.Next = &ListNode{Val: 8, Prev: headchild1}
	headchild1.Next.Next = &ListNode{Val: 9, Prev: headchild1.Next}
	headchild1.Next.Next.Next = &ListNode{Val: 10, Prev: headchild1.Next.Next}

	headchild2 := &ListNode{Val: 11}
	headchild2.Next = &ListNode{Val: 12, Prev: headchild2}

	headchild1.Next.Child = headchild2
	head.Next.Next.Child = headchild1

	r := flatten(head)
	re := []int{}

	for r != nil {
		re = append(re, r.Val)
		r = r.Next
	}
	t.Log(re)
}
