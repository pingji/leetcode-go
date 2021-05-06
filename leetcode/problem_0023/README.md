# Solutions
参考：[合并 K 个有序链表 - 4 种方法详解](https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/4-chong-fang-fa-xiang-jie-bi-xu-miao-dong-by-sweet/)

## Solution 1.1： K指针
K 个指针分别指向 K 条链表，每次 O(K) 比较 K个指针求 min

复杂度
- 时间复杂度：O(NK)

## Solution 1.2： K指针 + 最小堆
使用小根堆对 1 进行优化，每次 O(logK) 比较 K个指针求 min

复杂度
- 时间复杂度：O(NlogK)

## Solution 2.1： 逐一合并两条链表 

复杂度
- 时间复杂度：O(NK)

## Solution 2.2： 两两合并两条链表 

复杂度
- 时间复杂度：O(NlogK)

# 总结
掌握两种 O(NlogK) 方法：
- K 指针指向 K 条链表，每次使用小根堆来 logK 求出最小值；
- 对 K 条链表进行两两合并（递归 / 迭代）。
