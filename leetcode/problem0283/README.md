# [283. Move Zeroes](https://leetcode-cn.com/problems/move-zeroes/)
## Problem
```
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
```
## Solutions
### Solution 1: 一次遍历
思考快速排序：即选择一个数，比这个数小的放左边，比数大的放右边。
这里可以改变判断条件：不等于0的放左边，等于0的放右边
使用快慢指针，只要nums[i] != 0, 交换 nums[i] 和 nums[j]，同时移动指针

