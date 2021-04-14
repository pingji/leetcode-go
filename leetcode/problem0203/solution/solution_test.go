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
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	list := initList(arr)
	printList(list)
	list = removeElements(list, 6)
	printList(list) // [1,2,3,4,5]

	arr = []int{}
	list = initList(arr)
	printList(list)
	list = removeElements(list, 1)
	printList(list) // []

	arr = []int{7, 7, 7, 7}
	list = initList(arr)
	printList(list)
	list = removeElements(list, 7)
	printList(list) // []
}
