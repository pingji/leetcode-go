# [1025. Divisor Game](https://leetcode.com/problems/divisor-game/)
## Problem
```
Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number N on the chalkboard.  On each player's turn, that player makes a move consisting of:

Choosing any x with 0 < x < N and N % x == 0.
Replacing the number N on the chalkboard with N - x.
Also, if a player cannot make a move, they lose the game.

Return True if and only if Alice wins the game, assuming both players play optimally.

 

Example 1:

Input: 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.
Example 2:

Input: 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
 

Note:

1 <= N <= 1000.
```
## Solutions
### Solution 1: Dynamic Programming
为了便于计算，数组中需要添加最后到达的那一级台阶，cost为0。到达最后一级台阶就是到达终点。dp[i]为到达 第i级台阶需要的最小cost
```
dp[0] = cost[0]
dp[1] = cost[1]
dp[i] = min(dp[n-1], dp[n-2]) + cost[i]
```