package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		value := sum % 10
		carry = sum / 10
		curr.Next = &ListNode{value, nil}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}
