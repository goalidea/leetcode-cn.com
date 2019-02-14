package twolistintersection

import "testing"

func Test2ListIntersection(t *testing.T) {

	intersection := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: intersection,
		},
	}
	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: intersection,
			},
		},
	}

	r := getIntersectionNode(headA, headB)

	if r == intersection {
		t.Logf("Intersected at '%d'", r.Val)
	}
}
