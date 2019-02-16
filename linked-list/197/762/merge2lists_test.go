package merge2lists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	sl1, sl2 := []int{1, 2, 4}, []int{1, 3, 4}
	l1, l2 := &ListNode{}, &ListNode{}

	tmp1, tmp2 := l1, l2
	for i, v := range sl1 {
		tmp1.Val = v
		if i < len(sl1)-1 {
			tmp1.Next = &ListNode{}
			tmp1 = tmp1.Next
		}
	}

	for i, v := range sl2 {
		tmp2.Val = v
		if i < len(sl2)-1 {
			tmp2.Next = &ListNode{}
			tmp2 = tmp2.Next
		}
	}

	r := mergeTwoLists(l1, l2)

	re := []int{}

	for r != nil {
		re = append(re, r.Val)
		r = r.Next
	}

	t.Log(re)
}
