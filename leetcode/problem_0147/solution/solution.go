package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	for cur := head; cur != nil; {
		prev := dummy
		for prev.Next != nil {
			if cur.Val <= prev.Next.Val {
				break
			}
			prev = prev.Next
		}
		next := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return dummy.Next
}
