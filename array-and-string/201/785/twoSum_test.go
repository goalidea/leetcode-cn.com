// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-08 22:56
 */

package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	numbers, target, exprst := []int{2, 7, 11, 15}, 9, []int{1, 2}
	r := twoSum(numbers, target)
	for i, v := range r {
		if v != exprst[i] {
			t.Error("func twoSum error")
		}
	}
	t.Logf("input: numbers = %v target = %v", numbers, target)
	t.Logf("output: %v", r)
}
