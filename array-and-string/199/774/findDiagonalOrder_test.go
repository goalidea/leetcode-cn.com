// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-05 21:20
 */

package findDiagonalOrder

import "testing"

func TestFindDiagonalOrder(t *testing.T) {
	matrix, epre := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	r := findDiagonalOrder(matrix)
	for i, v := range r {
		if v != epre[i] {
			t.Error("func findDiagonalorder error")
		}
	}
	t.Log(r)
}
