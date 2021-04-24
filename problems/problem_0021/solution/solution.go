package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// dummyHead
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	prev := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}

	return dummyHead.Next
}
