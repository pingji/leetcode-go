# [303. Range Sum Query - Immutabler](https://leetcode-cn.com/problems/range-sum-query-immutable/)
## Problem
```
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
Note:
You may assume that the array does not change.
There are many calls to sumRange function.
```
## Solutions
### Solution 1: Dynamic Programming
由于sumRange 需要多次调用，最好的办法就是做一个缓存保存 sumRange(0,i)
```
sumRange(i,j) = sumRange(0,j) - sumRange(0.i-1)
```
另外，测试用例需要考虑 nums为nil的情况