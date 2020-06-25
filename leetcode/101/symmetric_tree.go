package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return same(root, root)
}

func same(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	// either is nil
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && same(t1.Right, t2.Left) && same(t1.Left, t2.Right)

}
