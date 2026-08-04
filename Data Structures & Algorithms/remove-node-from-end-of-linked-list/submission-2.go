/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}
	
	needFront :=  length - n 
	if needFront == 0 {
		head = head.Next
		return head 
	}
	//должен найти last front node
	lastFront := head
	for i := 1; i < needFront; i++ {
		lastFront = lastFront.Next
	}

	var recursion func(node *ListNode, depth int)
	recursion = func(node *ListNode, depth int) {
		if node == nil {
			return 
		}
		recursion(node.Next, depth+1)
		//идем обратно
		pos := length - depth + 1
		if pos == n {
			lastFront.Next = node.Next
		}
	}

	recursion(head, 1)
	return head
}
