// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 17:43
 */

package getRow

func getRow(rowIndex int) []int {
	rst := make([]int, rowIndex+1)

	rst[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			rst[j] += rst[j-1]
		}
	}
	return rst
}
