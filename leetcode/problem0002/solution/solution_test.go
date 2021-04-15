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
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	list1 := initList(arr1)
	list2 := initList(arr2)
	printList(list1)
	printList(list2)
	list3 := addTwoNumbers(list1, list2)
	printList(list3) // [7,0,8]

	arr1 = []int{0}
	arr2 = []int{0}
	list1 = initList(arr1)
	list2 = initList(arr2)
	printList(list1)
	printList(list2)
	list3 = addTwoNumbers(list1, list2)
	printList(list3) // [0]

	arr1 = []int{9, 9, 9, 9, 9, 9, 9}
	arr2 = []int{9, 9, 9, 9}
	list1 = initList(arr1)
	list2 = initList(arr2)
	printList(list1)
	printList(list2)
	list3 = addTwoNumbers(list1, list2)
	printList(list3) // [8,9,9,9,0,0,0,1]
}
