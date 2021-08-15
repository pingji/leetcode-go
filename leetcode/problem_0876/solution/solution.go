package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
