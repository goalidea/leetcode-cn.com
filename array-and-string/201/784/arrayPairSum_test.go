// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-08 22:41
 */

package arrayPairSum

import "testing"

func TestArrayPairSum(t *testing.T) {
	nums, exprst := []int{1, 4, 3, 2}, 4

	t.Logf("input: nums = %v", nums)
	rst := arrayPairSum(nums)
	if rst != exprst {
		t.Error("func arrayPairSum error")
	}
	t.Logf("output: %v", rst)
}
