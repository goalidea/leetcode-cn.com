// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-04 16:25
 */

package rotatelist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	//head长度是0或1的时候直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	//记录head长度变量count
	var count = 1
	var q *ListNode
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		count++
	}

	//设置为环list
	tmp.Next = head
	for i, n := 0, count-k%count; i < n; i++ {
		tmp = tmp.Next
	}
	q = tmp.Next
	tmp.Next = nil
	return q
}
