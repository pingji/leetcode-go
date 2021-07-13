package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func partition(head *ListNode, x int) *ListNode {
	dummySmall := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummyLarge := &ListNode{
		Val:  0,
		Next: nil,
	}
	small := dummySmall
	large := dummyLarge
	cur := head
	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			large.Next = cur
			large = large.Next
		}
		cur = cur.Next
	}
	large.Next = nil
	small.Next = dummyLarge.Next

	return dummySmall.Next
}
