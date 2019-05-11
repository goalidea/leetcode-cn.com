// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-11 18:10
 */

package reverseWords

import "strings"

func reverseWords(s string) string {
	sls := strings.Fields(s)
	for i := range sls {
		sls[i] = reverseStrings(sls[i])
	}
	return strings.Join(sls, " ")
}

func reverseStrings(s string) string {
	tmp, ls := []byte{}, len(s)
	for i := ls - 1; i >= 0; i-- {
		tmp = append(tmp, s[i])
	}
	return string(tmp)
}
