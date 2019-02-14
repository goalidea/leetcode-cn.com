package rmlistnode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	addnode := &ListNode{}
	addnode.Next = head

	tmp := addnode
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
			continue
		}
		tmp = tmp.Next
	}

	return addnode.Next
}
