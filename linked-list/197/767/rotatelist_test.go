// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-04 16:26
 */

package rotatelist

import "testing"

func TestRotateRight(t *testing.T) {
	head, k, testint, testres := &ListNode{}, 2, [5]int{1, 2, 3, 4, 5}, [5]int{4, 5, 1, 2, 3}
	tmp := head
	for _, v := range testint {
		tmp.Val = v
		if v == 5 {
			continue
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	r, re := rotateRight(head, k), []int{}
	for r != nil {
		re = append(re, r.Val)
		r = r.Next
	}
	for i, v := range testres {
		if re[i] != v {
			t.Error("func rotateRight error")
		}
	}
	t.Log(re)
}
