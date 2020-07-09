package main

import "fmt"

func main() {

	n := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(widthOfBinaryTree(n))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Index struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	storage := []Index{{node: root, index: 0}}

	for len(storage) > 0 {
		head := storage[0]
		size := len(storage)
		var last Index
		for i := 0; i < size; i++ {
			last = storage[0]
			storage = append(storage[:0], storage[1:]...) // remove head
			node := last.node
			if node.Left != nil {
				storage = append(storage, Index{node: node.Left, index: last.index * 2})
			}

			if node.Right != nil {
				storage = append(storage, Index{node: node.Right, index: last.index*2 + 1})
			}
		}

		result = max(result, last.index-head.index+1)
	}

	return result
}

func max(l, r int) int {
	if l > r {
		return l
	}

	return r
}
