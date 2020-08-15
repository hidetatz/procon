package main

/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {
	arr := []func(){}
	for {
		if head == nil {
			break
		}
		arr = append(arr, head.printValue)
		head = head.getNext()
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[i]()
	}
}
