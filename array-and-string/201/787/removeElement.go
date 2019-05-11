// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 12:17
 */

package removeElement

func removeElement(nums []int, val int) int {
	var k int
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}
