# [322. Coin Change](https://leetcode-cn.com/problems/coin-change/)
## Problem
```
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3 
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
```
## Solutions
### Solution 1: Dynamic Programming
dp[i]表示为组成金额 i 所需最少的硬币数量，转移方程为：
$$ dp(i) = \min_{j=0...n-1}dp(i−c_{j})+1 $$
其中 cj 代表的是第 j 枚硬币的面值