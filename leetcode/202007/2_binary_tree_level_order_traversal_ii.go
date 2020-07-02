package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrderBottom(r))
}

var result [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	result = [][]int{}
	if root == nil {
		return result
	}

	walk(root, 0)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func walk(n *TreeNode, level int) {
	if len(result) == level {
		result = append(result, []int{})
	}

	result[level] = append(result[level], n.Val)

	if n.Left != nil {
		walk(n.Left, level+1)
	}

	if n.Right != nil {
		walk(n.Right, level+1)
	}
}
