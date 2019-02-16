package merge2lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mergelist := &ListNode{}

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	tmp := mergelist
	for {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
			tmp = tmp.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
			tmp = tmp.Next
		}

		if l1 == nil {
			tmp.Next = l2
			break
		}
		if l2 == nil {
			tmp.Next = l1
			break
		}
	}
	return mergelist.Next
}
