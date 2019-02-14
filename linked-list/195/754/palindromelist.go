package palindromelist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	midnode := middlenode(head)
	midnode.Next = reverselist(midnode.Next)

	n1, n2 := head, midnode.Next
	for n1 != nil && n2 != nil && n1.Val == n2.Val {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n2 == nil
}

// find middle of node in list
func middlenode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// reverse list
func reverselist(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var nl *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = nl
		nl = head
		head = tmp
	}
	return nl
}
