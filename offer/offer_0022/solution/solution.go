package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	first, second := head, head
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}
