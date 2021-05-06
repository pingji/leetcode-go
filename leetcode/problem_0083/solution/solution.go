package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	curr := dummyHead

	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			val := curr.Next.Val
			curr = curr.Next
			for curr.Next != nil && curr.Next.Val == val {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}

	return dummyHead.Next
}
