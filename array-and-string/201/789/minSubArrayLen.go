// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 16:45
 */

package minSubArrayLen

func minSubArrayLen(s int, nums []int) int {
	lnums, minlen, left, right, sum := len(nums), len(nums)+10, 0, 0, 0

	for left < lnums {
		if right < lnums && sum < s {
			sum += nums[right]
			right++
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= s {
			if minlen > right-left {
				minlen = right - left
			}
		}
	}

	if minlen == lnums+10 {
		return 0
	}
	return minlen
}
