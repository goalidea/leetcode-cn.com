// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-11 18:39
 */

package moveZeroes

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums, exprst := []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}
	t.Logf("input: nums = %v", nums)
	moveZeroes(nums)
	for i, v := range nums {
		if v != exprst[i] {
			t.Error("func moveZeroes error")
		}
	}
	t.Logf("output: %v", nums)
}
