// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-11 16:25
 */

package reverseWords

import "testing"

func TestReverseWords(t *testing.T) {
	s, exprst := "the sky is blue", "blue is sky the"

	r := reverseWords(s)
	if r != exprst {
		t.Error("func reverseWords error")
	}
	t.Logf("input: s = %v", s)
	t.Logf("output: %v", r)
}
