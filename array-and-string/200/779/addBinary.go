// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-07 14:41
 */

package addBinary

import "fmt"

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	indexa, indexb, carry := la-1, lb-1, 0
	var rst string
	for indexa >= 0 || indexb >= 0 {
		tmpa, tmpb, tmp := 0, 0, 0
		if indexa >= 0 {
			tmpa = int(a[indexa] - '0')
		} else {
			tmpa = 0
		}

		if indexb >= 0 {
			tmpb = int(b[indexb] - '0')
		} else {
			tmpb = 0
		}

		tmp = tmpa + tmpb + carry
		rst = fmt.Sprintf("%d", tmp%2) + rst
		carry = tmp / 2
		indexa--
		indexb--
	}
	if carry == 1 {
		rst = "1" + rst
	}
	return rst
}
