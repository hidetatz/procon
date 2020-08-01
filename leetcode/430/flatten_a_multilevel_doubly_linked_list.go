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

func flatten(root *Node) *Node {
	p := root
	for p != nil {
		if p.Child != nil {
			pc := p.Child
			for pc.Next != nil {
				pc = pc.Next
			}
			pc.Next = p.Next
			if p.Next != nil {
				p.Next.Prev = pc
			}
			p.Next = p.Child
			p.Child.Prev = p
			p.Child = nil
		}
		p = p.Next
	}

	return root
}
