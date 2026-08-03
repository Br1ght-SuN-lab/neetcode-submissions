/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy 

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			tail = tail.Next
			list1 = list1.Next
		} else {
			tail.Next = list2
			tail = tail.Next
			list2 = list2.Next
		}
	}

	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}
	return dummy.Next
}
