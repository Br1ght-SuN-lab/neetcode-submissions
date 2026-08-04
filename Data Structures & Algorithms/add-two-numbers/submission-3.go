/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy 
	buf := 0

	for l1 != nil && l2 != nil {
		newNode := &ListNode{Val: (l1.Val + l2.Val) % 10 + buf}
		buf = (l1.Val + l2.Val + buf) / 10
		tail.Next = newNode
		tail = tail.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		for buf != 0 && l2 != nil {
			newNode := &ListNode{Val: (l2.Val + buf) % 10}
			tail.Next = newNode
			buf = (l2.Val + buf) / 10
			tail = tail.Next
			l2 = l2.Next
		}
		if l2 != nil {
			tail.Next = l2
		} else {
			if buf != 0 {
				tail.Next = &ListNode{Val: buf}
			} 
		}
	} else {
		for buf != 0 && l1 != nil {
			newNode := &ListNode{Val: (l1.Val + buf) % 10}
			tail.Next = newNode
			buf = (l1.Val + buf) / 10
			tail = tail.Next
			l1 = l1.Next
		}
		if l1 != nil {
			tail.Next = l1
		} else {
			if buf != 0 {
				tail.Next = &ListNode{Val: buf}
			} 
		}
	}

	return dummy.Next
}
