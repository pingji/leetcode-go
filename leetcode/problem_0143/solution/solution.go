package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	head1 := head
	head2 := mid.Next
	mid.Next = nil
	head2 = reverseList(head2)

	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	for head1 != nil || head2 != nil {
		if head1 != nil {
			dummy.Next = head1
			head1 = head1.Next
			dummy = dummy.Next
		}
		if head2 != nil {
			dummy.Next = head2
			head2 = head2.Next
			dummy = dummy.Next
		}
	}
	head = dummy.Next
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
