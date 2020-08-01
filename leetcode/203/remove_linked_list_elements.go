package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	node := &ListNode{}
	node.Next = head

	prev := node
	cur := head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next // remove
		} else {
			prev = cur
		}

		cur = cur.Next
	}

	return node.Next
}
