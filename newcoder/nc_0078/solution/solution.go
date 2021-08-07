package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	curr := pHead
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

	}
	return prev
}
