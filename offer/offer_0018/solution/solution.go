package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummyHead
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
	}

	return dummyHead.Next
}
