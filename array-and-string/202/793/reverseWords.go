// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-11 16:24
 */

package reverseWords

import "strings"

func reverseWords(s string) string {
	slices := strings.Fields(s)
	lslices := len(slices)
	for i, j := 0, lslices-1; i != j && i < j; {
		slices[i], slices[j] = slices[j], slices[i]
		i++
		j--
	}
	return strings.Join(slices, " ")
}
