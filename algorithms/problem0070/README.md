# [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)
## Problem
```
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 stepp
```
## Solutions
### Solution 1: Dynamic Programming
第 ii 阶可以由以下两种方法得到：
- 在第 (i-1)阶后向上爬1阶。
- 在第 (i-2)阶后向上爬2阶。
```
dp[0] = 1
dp[1] = 2
dp[i] = dp[i-2] +dp[i-1]
```
