package solution

import (
	"container/heap"
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	// dummyHead
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	prevNode := dummyHead

	var h NodeHeap
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&h, lists[i])
		}
	}

	for h.Len() > 0 {
		currNode := heap.Pop(&h).(*ListNode)
		nextNode := currNode.Next
		if nextNode != nil {
			heap.Push(&h, nextNode)
		}

		currNode.Next = nil
		prevNode.Next = currNode
		prevNode = prevNode.Next
	}

	return dummyHead.Next
}
