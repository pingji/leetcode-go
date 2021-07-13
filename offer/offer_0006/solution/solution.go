package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

func reversePrint(head *ListNode) []int {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	res := make([]int, count)

	cur = head
	index := count - 1
	for cur != nil {
		res[index] = cur.Val
		cur = cur.Next
		index--
	}

	return res
}
