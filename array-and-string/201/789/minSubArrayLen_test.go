// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 16:46
 */

package minSubArrayLen

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	s, nums, exprst := 7, []int{2, 3, 1, 2, 4, 3}, 2
	r := minSubArrayLen(s, nums)
	if r != exprst {
		t.Error("func minSubArrayLen error")
	}
	t.Logf("input: s = %v nums = %v", s, nums)
	t.Logf("output: %v", r)
}
