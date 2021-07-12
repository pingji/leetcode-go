package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummyHead
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := pre.Next.Next
		pre.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		pre = node1
	}

	return dummyHead.Next
}
