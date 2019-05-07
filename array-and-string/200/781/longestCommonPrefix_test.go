// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-07 19:02
 */

package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs, exprst := []string{"flower", "flow", "flight"}, "fl"
	r := longestCommonPrefix(strs)
	if r != exprst {
		t.Error("func longestCommonPrefix error")
	}
	t.Logf("input: strs = %v", strs)
	t.Logf("output: %v", r)
}
