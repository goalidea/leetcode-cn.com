// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-07 14:42
 */

package addBinary

import "testing"

func TestAddBinary(t *testing.T) {
	a, b, expres := "11", "1", "100"
	r := addBinary(a, b)
	if r != expres {
		t.Error("func addBinary error")
	}
	t.Logf("input: a = %v b = %v", a, b)
	t.Logf("output: %v", r)
}
