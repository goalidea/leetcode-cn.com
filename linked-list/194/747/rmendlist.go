package rmendlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nrd, end := head, head
	for i := 0; i < n; i++ {
		end = end.Next
	}
	if end == nil {
		return head.Next
	}
	for end.Next != nil {
		nrd = nrd.Next
		end = end.Next
	}
	nrd.Next = nrd.Next.Next
	return head
}
