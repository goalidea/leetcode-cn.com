// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-08 22:04
 */

package reverseString

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	t.Logf("input: s = %c", s)
	reverseString(s)
	t.Logf("change: s = %c", s)
}
