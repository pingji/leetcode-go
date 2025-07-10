package solution

import (
	"leetcode-go/common/linked_list"
)

type ListNode = linked_list.ListNode

// mergeSort 合并两个有序链表
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func mergeSort(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建虚拟头节点，简化合并过程
	dummy := &ListNode{Val: 0, Next: nil}
	prev := dummy

	// 比较两个链表的节点，选择较小的节点添加到结果链表
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	// 将剩余节点添加到结果链表
	if l1 != nil {
		prev.Next = l1
	} else if l2 != nil {
		prev.Next = l2
	}

	return dummy.Next
}

// sortList 对链表进行归并排序
// 时间复杂度: O(n log n)
// 空间复杂度: O(log n) - 递归调用栈深度
func sortList(head *ListNode) *ListNode {
	// 递归终止条件：空链表或单节点链表
	if head == nil || head.Next == nil {
		return head
	}

	// 使用虚拟头节点和快慢指针找到链表的中点
	// 虚拟头节点确保快慢指针从同一位置开始，避免边界情况
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将链表分割为两个子链表
	// slow.Next 是右半部分的起始节点
	next := slow.Next
	slow.Next = nil // 断开链表，形成两个独立的子链表

	// 递归排序两个子链表
	l := sortList(head) // 排序左半部分
	r := sortList(next) // 排序右半部分

	// 合并两个有序链表
	return mergeSort(l, r)
}
