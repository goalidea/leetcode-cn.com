package reversesinglylist

import "testing"

func TestReverseSinglyList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	r := reverseList(head)
	re := []int{}
	for r != nil {
		re = append(re, r.Val)
		r = r.Next
	}

	t.Log(re)
}
