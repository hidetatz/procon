package main

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	var p, q, curr = l1, l2, res
	var carry = 0

	for p != nil || q != nil {
		sum := carry
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		fmt.Printf("sum = %+v\n", sum)
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return res.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l := addTwoNumbers(l1, l2)
	for {
		fmt.Printf("%+v\n", l)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
}
