// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-06 19:53
 */

package spiralOrder

import "testing"

func TestSpiralOrder(t *testing.T) {
	matrix, expresult := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	r := spiralOrder(matrix)
	for i, v := range r {
		if v != expresult[i] {
			t.Error("func spiralOrder error")
		}
	}
	t.Logf("matrix is %v", matrix)
	t.Logf("result is %v", r)
}
