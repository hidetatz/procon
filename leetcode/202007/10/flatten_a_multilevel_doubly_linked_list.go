package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

var last *Node

func flatten(root *Node) *Node {
	helper(root)
	return root
}

func helper(n *Node) {
	if n != nil {
		if last != nil {
			n.Prev = last
			last.Next = n
		}

		last = n

		newNode := n.Next

		helper(n.Child)

		n.Child = nil

		helper(newNode)
	}
}
