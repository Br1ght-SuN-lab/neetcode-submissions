/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func recursion(node *ListNode, depth int) {
	if (node == nil) {
		return 
	}
	recursion(node.Next, depth+1)
	//здесь будет разворот рекурсии

}

func reorderList(head *ListNode) {
	tail := head
	frontPtr := head.Next

	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	half := n / 2
	needFront := n - 1 - half //в tail лежит уже один элемент из Front

	var recursion func(node *ListNode, depth int)
	recursion = func(node *ListNode, depth int) {
		if (node == nil) {
			return 
		}
		recursion(node.Next, depth+1)
		//здесь будет разворот рекурсии (node разворачивается)
		posFromEnd := n - depth + 1 
		
		if posFromEnd <= half {
			tail.Next = node 
			tail = node
			if posFromEnd <= needFront {
				tail.Next = frontPtr
				tail = frontPtr
				frontPtr = frontPtr.Next
			}
		}
	}

	recursion(head, 1)
	tail.Next = nil
}
