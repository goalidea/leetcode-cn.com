package oddevenlist

import "testing"

func TestOddEvenList(t *testing.T) {
	head := &ListNode{}
	val := []int{1, 2, 3, 4, 5}

	tmp := head
	for i, v := range val {
		tmp.Val = v
		if i == len(val)-1 {
			break
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}

	r := oddEvenList(head)
	re := []int{}
	for r != nil {
		re = append(re, r.Val)
		r = r.Next
	}

	t.Log(re)
}
