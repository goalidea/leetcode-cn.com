package reversesinglylist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	firstnode, nextnode := head, head.Next
	for nextnode != nil {
		head.Next = nextnode.Next
		nextnode.Next = firstnode
		firstnode = nextnode
		nextnode = head.Next
	}
	return firstnode
}
