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

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	for step := 1; step < length; step += step {
		prev := dummyHead
		cur := dummyHead.Next
		for cur != nil {
			leftHead := cur
			for i := 1; i < step && cur.Next != nil; i++ {
				cur = cur.Next
			}

			if cur.Next == nil {
				prev.Next = leftHead
				cur = nil
			} else {
				rightHead := cur.Next
				cur.Next = nil
				cur = rightHead
				for i := 1; i < step && cur != nil && cur.Next != nil; i++ {
					cur = cur.Next
				}
				var next *ListNode
				if cur != nil {
					next = cur.Next
					cur.Next = nil
				}

				prev.Next = mergeSort(leftHead, rightHead)
				for prev.Next != nil {
					prev = prev.Next
				}
				cur = next
			}
		}
	}
	return dummyHead.Next
}
