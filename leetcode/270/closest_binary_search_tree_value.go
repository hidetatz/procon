package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	ret := root.Val
	del := delta(root.Val, target)

	var cur *TreeNode
	if float64(root.Val) < target {
		if root.Right == nil {
			return ret
		} else {
			cur = root.Right
		}
	} else {
		if root.Left == nil {
			return ret
		} else {
			cur = root.Left
		}
	}

	for {
		d := delta(cur.Val, target)
		if del > d {
			ret = cur.Val
			del = d
		}

		if float64(cur.Val) < target {
			if cur.Right == nil {
				break
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Left == nil {
				break
			} else {
				cur = cur.Left
			}
		}
	}

	return ret
}

func delta(i int, target float64) float64 {
	if float64(i) > target {
		return float64(i) - target
	}
	return target - float64(i)
}
