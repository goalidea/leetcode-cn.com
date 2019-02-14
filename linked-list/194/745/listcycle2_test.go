package listcycle2

import "testing"

func TestListCycle2(t *testing.T) {
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

	r := detectCycle(head)

	if r == head.Next {
		t.Log("tail connects to node index 1")
	}
}
