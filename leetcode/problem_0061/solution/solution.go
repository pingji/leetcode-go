package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	n := 1
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	step := n - k%n
	if step == n {
		return head
	}
	cur.Next = head
	for i := 0; i < step; i++ {
		cur = cur.Next
	}
	ret := cur.Next
	cur.Next = nil
	return ret
}
