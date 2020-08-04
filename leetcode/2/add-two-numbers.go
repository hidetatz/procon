package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	ret := &ListNode{Next: cur}

	carry := 0
	for (l1 != nil || l2 != nil) || carry != 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			sum = sum - 10
		}

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	return ret.Next.Next
}
