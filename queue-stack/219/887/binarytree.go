package binarytree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	return dfs(root)
}

func dfs(root *TreeNode) []int {
	result := []int{}

	stack, node := list.New(), root

	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

		if stack.Len() > 0 {
			tmp := stack.Back().Value.(*TreeNode)
			stack.Remove(stack.Back())
			result = append(result, tmp.Val)
			node = tmp.Right
		}
	}

	return result
}
