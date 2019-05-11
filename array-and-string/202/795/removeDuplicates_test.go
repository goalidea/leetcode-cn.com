// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-11 18:32
 */

package removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums, exprst := []int{1, 1, 2}, 2
	t.Logf("input: nums = %v", nums)
	r := removeDuplicates(nums)
	if r != exprst {
		t.Error("func removeDuplicates error")
	}
	t.Logf("output: %v", nums[:r])
}
