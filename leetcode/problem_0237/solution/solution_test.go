package solution

import (
	"fmt"
	"testing"
)

func initList(arr []int) *ListNode {
	// dummyHead
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dummyHead
	for _, value := range arr {
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		cur.Next = node
		cur = node
	}
	return dummyHead.Next
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%v=>", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("NULL\n")
}

func foundNode(head *ListNode, index int) *ListNode {
	curr := head
	for i := 0; curr != nil; i++ {
		if i == index {
			return curr
		}
		curr = curr.Next
	}
	return nil
}
func TestProblem(t *testing.T) {
	arr := []int{4, 5, 1, 9}
	list := initList(arr)
	printList(list)
	node := foundNode(list, 1)
	deleteNode(node)
	printList(list) // [4,1,9]

	arr = []int{4, 5, 1, 9}
	list = initList(arr)
	printList(list)
	node = foundNode(list, 2)
	deleteNode(node)
	printList(list) // [4,5,9]
}
