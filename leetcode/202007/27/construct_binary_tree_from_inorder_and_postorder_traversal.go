package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	m := map[int]int{}
	for i, v := range inorder {
		m[v] = i
	}

	var helper func(int, int) *TreeNode
	idx := len(postorder) - 1

	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := &TreeNode{Val: postorder[idx]}
		idx--
		v := m[root.Val]
		root.Right = helper(v+1, right)
		root.Left = helper(left, v-1)
		return root
	}

	return helper(0, len(inorder)-1)
}
