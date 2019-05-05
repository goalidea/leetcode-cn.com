// Copyright 2019 The Goalidea Authors. All rights reserved.
// Use of this source code is governed by a MPL-style
// license that can be found in the LICENSE file.

/**
 * @author goalidea<goalidea@outlook.com>
 * @create 2019-05-04 18:34
 */

package dominantIndex

import "testing"

func TestDominantIndex(t *testing.T) {
	nums, k := []int{0, 0, 0, 1}, 3
	r := dominantIndex(nums)
	if r != k {
		t.Error("func dominantIndex error")
	}
	t.Log(r)
}
