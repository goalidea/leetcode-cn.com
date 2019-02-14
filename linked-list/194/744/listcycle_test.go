package listcycle

import "testing"

func TestListCycle(t *testing.T) {

	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: -4,
				},
			},
		},
	}
	head.Next.Next.Next = head.Next

	r := hasCycle(head)

	if !r {
		t.Error("has cycle error")
	}
}
