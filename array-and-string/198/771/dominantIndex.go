// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-04 18:33
 */

package dominantIndex

func dominantIndex(nums []int) int {
	lnums := len(nums)
	//当nums为空是返回-1，nums只有一个元素的时候返回0
	if lnums == 0 {
		return -1
	} else if lnums == 1 {
		return 0
	}

	var firstMax, secondMax, k int
	//当nums有2个以上元素时，先比较第一个元素和第二个元素并复制给firstMax，secondMax变量
	if nums[0] > nums[1] {
		firstMax = nums[0]
		secondMax = nums[1]
	} else {
		firstMax = nums[1]
		secondMax = nums[0]
		k = 1
	}
	//从nums第三个元素开始遍历，找到最大数和第二大数并复制替换给firstMax,secondMax变量
	for i := 2; i < lnums; i++ {
		if firstMax < nums[i] {
			secondMax = firstMax
			firstMax = nums[i]
			k = i
		} else if secondMax < nums[i] {
			secondMax = nums[i]
		}
	}
	if firstMax >= 2*secondMax {
		return k
	}
	return -1
}
