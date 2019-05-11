// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-08 22:40
 */

package arrayPairSum

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums) // nums顺序排序

	var rst int
	for i, v := range nums {
		if i%2 == 0 {
			rst += v
		}
	}
	return rst
}
