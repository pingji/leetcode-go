package solution2

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
	arr := []int{1, 2, 3, 4, 5}
	list := initList(arr)
	printList(list)
	list = reverseBetween(list, 2, 4)
	printList(list) // [1,4,3,2,5]

	arr = []int{5}
	list = initList(arr)
	printList(list)
	list = reverseBetween(list, 1, 1)
	printList(list) // [5]

	arr = []int{3, 5}
	list = initList(arr)
	printList(list)
	list = reverseBetween(list, 1, 2)
	printList(list) // [5,3]
}
