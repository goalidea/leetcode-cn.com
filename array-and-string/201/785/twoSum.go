// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-08 22:56
 */

package twoSum

func twoSum(numbers []int, target int) []int {
	lnumbers := len(numbers)

	rst := []int{}
	for start, end := 0, lnumbers-1; start != end; {
		if numbers[start]+numbers[end] == target {
			rst = append(rst, start+1, end+1)
			break
		}
		if numbers[start]+numbers[end] > target {
			end--
			continue
		}
		if numbers[start]+numbers[end] < target {
			start++
			continue
		}
	}
	return rst
}
