// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-06 19:52
 */

package spiralOrder

func spiralOrder(matrix [][]int) []int {
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
	for r, c := 0, 0; m > 0 && n > 0; {
		//最后只剩一行
		if m == 1 {
			for i := 0; i < n; i++ {
				result = append(result, matrix[r][c])
				c++
			}
			break
		}
		//最后只剩一列
		if n == 1 {
			for i := 0; i < m; i++ {
				result = append(result, matrix[r][c])
				r++
			}
			break
		}

		//向右遍历
		for i := 0; i < n-1; i++ {
			result = append(result, matrix[r][c])
			c++
		}
		//向下遍历
		for i := 0; i < m-1; i++ {
			result = append(result, matrix[r][c])
			r++
		}
		//向左遍历
		for i := 0; i < n-1; i++ {
			result = append(result, matrix[r][c])
			c--
		}
		//向上遍历
		for i := 0; i < m-1; i++ {
			result = append(result, matrix[r][c])
			r--
		}
		r++
		c++
		m -= 2
		n -= 2
	}
	return result
}
