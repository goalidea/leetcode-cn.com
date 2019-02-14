package binarytree

import (
	"testing"
)

func TestBinarytree(t *testing.T) {
	//[1,null,2,3]
	root := &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 2,
	}
	root.Right.Left = &TreeNode{
		Val: 3,
	}
	expectation := []int{1, 3, 2}

	r := inorderTraversal(root)

	for i, v := range r {
		if v != expectation[i] {
			t.Error("traversal binarytree error")
		}
	}
}
