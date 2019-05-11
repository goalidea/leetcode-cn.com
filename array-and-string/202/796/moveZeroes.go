// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-11 18:39
 */

package moveZeroes

func moveZeroes(nums []int) {
	lnums := len(nums)
	for i, j := 0, lnums-1; i <= j; i++ {
		if nums[i] == 0 {
			for k := i; k < j; k++ {
				nums[k] = nums[k+1]
			}
			nums[j] = 0
			j--
			i--
		}
	}
}
