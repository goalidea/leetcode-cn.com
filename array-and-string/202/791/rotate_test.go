// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 17:31
 */

package rotate

import "testing"

func TestRotate(t *testing.T) {
	nums, k, exprst := []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}
	t.Logf("input: nums = %v k = %v", nums, k)
	rotate(nums, k)
	for i, v := range nums {
		if v != exprst[i] {
			t.Error("func rotate error")
		}
	}
	t.Logf("output: %v", nums)
}
