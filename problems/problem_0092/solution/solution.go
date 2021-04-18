package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var prev *ListNode
	var prevTail, currHead *ListNode
	var currTail, nextHead *ListNode
	curr := head
	index := 1
	for curr != nil && index <= right {
		next := curr.Next
		if index >= left && index <= right {
			if index == left {
				prevTail = prev
				currHead = curr
			}
			if index == right {
				currTail = curr
				nextHead = next
			}
			curr.Next = prev
		}
		prev = curr
		curr = next
		index++
	}
	if prevTail != nil {
		prevTail.Next = currTail
	} else {
		head = currTail
	}
	currHead.Next = nextHead

	return head
}
