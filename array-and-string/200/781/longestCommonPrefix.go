// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-07 19:02
 */

package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	lstrs := len(strs)
	if strs == nil {
		return ""
	}
	if lstrs == 0 {
		return ""
	} else if lstrs == 1 {
		return strs[0]
	}

	lstrsIndex1, rst := len(strs[0]), []byte{}
ooo:
	for i := 0; i < lstrsIndex1; i++ {
		for j := 1; j < lstrs; j++ {
			if strs[j] == "" {
				return ""
			}
			if i == 0 && strs[0][i] != strs[j][i] {
				return ""
			}

			if i < len(strs[j]) && strs[0][i] != strs[j][i] {
				break ooo
			}
			if i >= len(strs[j]) {
				break ooo
			}
		}
		rst = append(rst, strs[0][i])
	}
	return string(rst)
}
