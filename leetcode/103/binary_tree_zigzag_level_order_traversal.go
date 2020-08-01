package main

var res [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level >= len(res) {
		res = append(res, []int{})
	}
	if level%2 == 0 { // left to right
		res[level] = append(res[level], root.Val)

	} else { // right to left
		temp := make([]int, len(res[level])+1)
		temp[0] = root.Val
		copy(temp[1:], res[level])
		res[level] = temp
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
