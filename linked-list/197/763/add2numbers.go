package add2numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumlist := &ListNode{}

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tmp1, tmp2 := l1, l2
	for {
		if tmp1.Next != nil && tmp2.Next != nil {
			tmp1 = tmp1.Next
			tmp2 = tmp2.Next
		}
		if tmp1.Next == nil && tmp2.Next != nil {
			tmp1.Next = &ListNode{}
			tmp1 = tmp1.Next
			tmp2 = tmp2.Next
		}
		if tmp1.Next != nil && tmp2.Next == nil {
			tmp1 = tmp1.Next
			tmp2.Next = &ListNode{}
			tmp2 = tmp2.Next
		}
		if tmp1.Next == nil && tmp2.Next == nil {
			break
		}
	}

	tmp := sumlist
	for {
		if tmp.Val+l1.Val+l2.Val > 9 {
			tmp.Val = (tmp.Val + l1.Val + l2.Val) % 10
			l1, l2 = l1.Next, l2.Next
			tmp.Next = &ListNode{
				Val: 1,
			}
			tmp = tmp.Next
		} else {
			tmp.Val = tmp.Val + l1.Val + l2.Val
			l1, l2 = l1.Next, l2.Next
			if l1 != nil && l2 != nil {
				tmp.Next = &ListNode{}
				tmp = tmp.Next
			}
		}

		if l1 == nil && l2 == nil {
			break
		}
	}

	return sumlist
}
