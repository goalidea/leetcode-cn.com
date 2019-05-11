// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 12:35
 */

package findMaxConsecutiveOnes

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums, exprst := []int{1, 0, 1, 1, 0, 1}, 2

	r := findMaxConsecutiveOnes(nums)
	if r != exprst {
		t.Error("func findMaxConsecutiveOnes error")
	}
	t.Logf("input: nums = %v", nums)
	t.Logf("output: %v", r)
}
