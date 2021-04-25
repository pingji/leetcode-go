package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// dummyHead
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	prev := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
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
	}
	if l2 != nil {
		prev.Next = l2
	}

	return dummyHead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	k := len(lists)
	for k > 1 {
		index := 0
		for i := 0; i < k; i = i + 2 {
			if i == k-1 {
				lists[index] = lists[i]
				index++
			} else {
				lists[index] = mergeTwoLists(lists[i], lists[i+1])
				index++
			}
		}
		k = index
	}

	return lists[0]
}
