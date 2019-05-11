// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 17:31
 */

package rotate

func rotate(nums []int, k int) {
	var tmp int
	l := len(nums)
	for ; k > 0; k-- {
		tmp = nums[l-1]
		for i := l - 1; i > 0; i-- {
			nums[i] = nums[i-1]
			if i == 1 {
				nums[0] = tmp
			}
		}
	}
}
