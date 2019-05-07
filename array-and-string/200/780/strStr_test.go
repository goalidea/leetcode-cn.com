// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-07 18:46
 */

package strStr

import "testing"

func TestStrStr(t *testing.T) {
	haystack, needle, exprst := "hello", "ll", 2
	r := strStr(haystack, needle)
	if r != exprst {
		t.Error("func strStr error")
	}
	t.Logf("input: haystack = %v needle = %v", haystack, needle)
	t.Logf("output: %v", r)
}
