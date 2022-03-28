package solution2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy

	index := 1
	var prevLeftNode *ListNode
	var leftNode *ListNode
	var rightNode *ListNode
	var nextRightNode *ListNode
	for prev.Next != nil && index <= right {
		if index == left {
			prevLeftNode = prev
			leftNode = prev.Next
		}
		if index == right {
			rightNode = prev.Next
			nextRightNode = prev.Next.Next
		}
		prev = prev.Next
		index++
	}

	leftNode, rightNode = reverse(leftNode, rightNode)
	prevLeftNode.Next = leftNode
	rightNode.Next = nextRightNode
	return dummy.Next
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode = nil
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next

	}
	return tail, head
}
