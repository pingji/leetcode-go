package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func mergeSort(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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
	} else if l2 != nil {
		prev.Next = l2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var prevSlow *ListNode
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prevSlow.Next = nil
	l := sortList(head)
	r := sortList(slow)
	return mergeSort(l, r)
}
