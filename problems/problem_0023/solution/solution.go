package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	// dummyHead
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	// prev := dummyHead

	return dummyHead.Next
}
