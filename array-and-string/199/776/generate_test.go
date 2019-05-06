// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-06 22:18
 */

package generate

import "testing"

func TestGenerate(t *testing.T) {
	numRows, expres := 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	r := generate(numRows)
	for i, iv := range r {
		for j, v := range iv {
			if v != expres[i][j] {
				t.Error("func generate error")
			}
		}
	}
	t.Logf("input: %v", numRows)
	t.Logf("output: %v", r)
}
