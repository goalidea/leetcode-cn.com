// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 12:17
 */

package removeElement

import "testing"

func TestRemoveElement(t *testing.T) {
	nums, val, exprst := []int{3, 2, 2, 3}, 3, 2
	t.Logf("input: nums = %v var = %v", nums, val)
	r := removeElement(nums, val)
	if r != exprst {
		t.Error("func removeElement error")
	}
	t.Logf("output: %v", r)
	t.Logf("change: nums = %v and front of nums = %v", nums, nums[:r])
}
