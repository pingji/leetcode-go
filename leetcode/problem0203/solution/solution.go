package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dummyHead
	curr := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = prev.Next
	}
	return dummyHead.Next
}
