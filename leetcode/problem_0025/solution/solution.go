package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode = nil
	cur := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// dummyHead
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummyHead

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
	}

	return dummyHead.Next
}
