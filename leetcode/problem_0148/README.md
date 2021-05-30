# Solutions

你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

数组额外空间：链表可以通过修改引用来更改节点顺序，无需像数组一样开辟额外空间；



## Solution 1：归并排序（递归）
### 二分
链表找到中心节点的方法：快慢指针法，快指针走两步，慢指针走一步，快指针越界时，慢指针正好到达中点，只需记录慢指针的前一个指针，就能断成两链。
### merge
设置虚拟头结点，用 prev 指针去“穿针引线”，
### 参考
- [手画图解归并排序](https://leetcode-cn.com/problems/sort-list/solution/shou-hua-tu-jie-gui-bing-pai-xu-148-lian-biao-pai-/)

## Solution 2：归并排序（迭代）

