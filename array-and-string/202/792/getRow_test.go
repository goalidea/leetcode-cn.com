// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-10 17:43
 */

package getRow

import "testing"

func TestGetRow(t *testing.T) {
	rowIndex, exprst := 3, []int{1, 3, 3, 1}

	r := getRow(rowIndex)
	for i, v := range r {
		if v != exprst[i] {
			t.Error("func getRow error")
		}
	}
	t.Logf("input: rowIndex = %v", rowIndex)
	t.Logf("output: %v", r)
}
