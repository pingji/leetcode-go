package linked_list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitList(arr []int) *ListNode {
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
func (head *ListNode) Convert2Slice() []int {
	res := []int{}
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func (head *ListNode) ToString() string {
	var b strings.Builder
	cur := head
	for cur != nil {
		fmt.Fprintf(&b, "%v=>", cur.Val)
		cur = cur.Next
	}
	fmt.Fprint(&b, "NULL")
	return b.String()
}

func (head *ListNode) PrintList() {
	cur := head
	for cur != nil {
		fmt.Printf("%v=>", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("NULL\n")
}

func (head *ListNode) IsEqual(target *ListNode) bool {
	cur1, cur2 := head, target
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil && cur2 == nil {
		return true
	}
	return false
}
