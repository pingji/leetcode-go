# [53. Maximum Subarray](https://leetcode-cn.com/problems/maximum-subarray/)
## Problem
```
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
```
## Solutions
### Solution 1: Dynamic Programming
若前一个元素的sum 大于0，则将它加到当前元素上
sum[0] = arr[0]
sum[i] = max(sum[i-1] +arr[i], arr[i])

if sum[i-1] > 0 : sum[i-1] +arr[i]
else arr[i]

### Solution 2: Greedy Algorithm
局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。