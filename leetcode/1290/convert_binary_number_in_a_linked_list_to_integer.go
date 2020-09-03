package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	ret := 0

	for head != nil {
		ret = ret*2 + head.Val
		head = head.Next
	}

	return ret
}
