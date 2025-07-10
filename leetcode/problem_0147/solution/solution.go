package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

// insertionSortList 对链表进行插入排序
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
func insertionSortList(head *ListNode) *ListNode {
	// 创建虚拟头节点，简化边界情况处理
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	// 遍历原链表的每个节点
	for cur := head; cur != nil; {
		// 在已排序的链表中找到插入位置
		prev := dummy
		for prev.Next != nil {
			// 找到第一个大于等于当前节点值的位置
			if cur.Val <= prev.Next.Val {
				break
			}
			prev = prev.Next
		}

		// 保存当前节点的下一个节点，避免丢失
		next := cur.Next
		// 将当前节点插入到正确位置
		cur.Next = prev.Next
		prev.Next = cur
		// 移动到下一个待排序的节点
		cur = next
	}

	// 返回排序后的链表（跳过虚拟头节点）
	return dummy.Next
}
