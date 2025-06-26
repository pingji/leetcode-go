# [739. Daily Temperatures](https://leetcode.cn/problems/daily-temperatures/)

## Solution 1

单调栈（Monotonic Stack）是一种特殊的栈结构，其内部元素保持单调递增或递减。它在算法问题中常用于高效解决区间极值和最近邻比较问题。

- map 里存放第一个数组的 value => index 的映射
- stack 里存放第二个数组的index，按照值的递增栈。遍历num2，
  - 如果当前值比栈顶元素小，则入栈
  - 如果当前值比栈顶元素大，则栈顶元素出栈，循环出栈，直到新的栈顶元素大于当前值或者栈为空
