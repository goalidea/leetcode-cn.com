// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-05 21:19
 */

package findDiagonalOrder

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	m := len(matrix)
	if m == 0 {
		return nil
	} else if m == 1 {
		return matrix[0]
	}

	n, result := len(matrix[0]), []int{}
	lens := m * n
	for i, r, c := 0, 0, 0; i < lens; i++ {
		result = append(result, matrix[r][c])
		if (r+c)%2 == 0 {
			if r == 0 && c != n-1 {
				c++
			} else if c == n-1 {
				r++
			} else {
				r--
				c++
			}
		} else {
			if c == 0 && r != m-1 {
				r++
			} else if r == m-1 {
				c++
			} else {
				r++
				c--
			}
		}
	}
	return result
}
