/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}

		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}

	return dummy.Next
}


func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	var result *ListNode = nil 
	for i := 0; i < n; i++ {
		list := lists[i]
		result = mergeTwoLists(result, list)
	}

	return result
}
