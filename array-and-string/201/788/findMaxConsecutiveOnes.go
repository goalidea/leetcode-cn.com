// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 12:35
 */

package findMaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	var count, max int

	for _, v := range nums {
		if v == 1 {
			count++
		}
		if v == 0 {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	return max
}
