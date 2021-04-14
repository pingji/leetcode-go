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
func TestProblem(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 4, 5}
	list := initList(arr)
	printList(list)
	list = deleteDuplicates(list)
	printList(list) // [1,2,5]

	arr = []int{1, 1, 1, 2, 3}
	list = initList(arr)
	printList(list)
	list = deleteDuplicates(list)
	printList(list) // [2,3]
}
