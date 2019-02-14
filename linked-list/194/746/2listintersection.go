package twolistintersection

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	// count len of headA and headB
	pointerA, pointerB := headA, headB
	var lA, lB int
	for pointerA != nil {
		lA++
		pointerA = pointerA.Next
	}
	for pointerB != nil {
		lB++
		pointerB = pointerB.Next
	}

	pointerA, pointerB = headA, headB
	if lA > lB {
		for i := 0; i < lA-lB; i++ {
			pointerA = pointerA.Next
		}
		for pointerA != nil {
			if pointerA == pointerB {
				return pointerA
			}
			pointerA = pointerA.Next
			pointerB = pointerB.Next
		}
		return nil
	}

	for i := 0; i < lB-lA; i++ {
		pointerB = pointerB.Next
	}

	for pointerB != nil {
		if pointerA == pointerB {
			return pointerB
		}
		pointerA = pointerA.Next
		pointerB = pointerB.Next
	}
	return nil
}
