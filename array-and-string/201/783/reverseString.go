// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-08 22:04
 */

package reverseString

func reverseString(s []byte) {
	ls := len(s)
	for i := 0; i < ls/2; i++ {
		s[i], s[ls-1-i] = s[ls-1-i], s[i]
	}
	return
}
