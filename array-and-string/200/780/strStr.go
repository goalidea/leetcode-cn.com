// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-07 18:46
 */

package strStr

import "strings"

func strStr(haystack string, needle string) int {
	//内置strings包
	if !strings.Contains(haystack, needle) {
		return -1
	}

	return strings.Index(haystack, needle)
}
