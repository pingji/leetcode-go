# [739. Daily Temperatures](https://leetcode.cn/problems/daily-temperatures/)

## Solution 1
单调栈（Monotonic Stack）是一种特殊的栈结构，其内部元素保持单调递增或递减。它在算法问题中常用于高效解决区间极值和最近邻比较问题。

单调栈里存储的是已经遍历过的元素的下标，单调递增

- s[i] > stack[top]: 计算天数，top元素出栈，继续查看下一个top元素，最后当前元素s[i]入栈
- s[i] <= stack[top]: 当前元素s[i]入栈