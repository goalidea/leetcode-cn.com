package rmlistnode

import "testing"

func TestRmListNode(t *testing.T) {
	val := 6
	sliceval := []int{1, 2, 6, 3, 4, 5, 6}
	head := &ListNode{}
	tmp := head
	for i, v := range sliceval {
		tmp.Val = v
		if i == len(sliceval)-1 {
			break
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	r := removeElements(head, val)
	re := []int{}
	for r != nil {
		re = append(re, r.Val)
		r = r.Next
	}
	t.Log(re)
}
