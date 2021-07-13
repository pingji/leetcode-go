package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	m := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		node := &Node{
			Val:    cur.Val,
			Next:   nil,
			Random: nil,
		}
		m[cur] = node
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		node := m[cur]
		node.Next = m[cur.Next]
		node.Random = m[cur.Random]
		cur = cur.Next
	}

	return m[head]
}
