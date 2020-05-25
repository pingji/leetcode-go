# [392. Is Subsequence](https://leetcode-cn.com/problems/is-subsequence/)
## Problem
```
Given a string s and a string t, check if s is subsequence of t.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).

Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?

Credits:
Special thanks to @pbrother for adding this problem and creating all test cases.

 

Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false
 

Constraints:

0 <= s.length <= 100
0 <= t.length <= 10^4
Both strings consists only of lowercase characters.
```
## Solutions
### Solution 1: Dynamic Programming
解题思路
1. 两个指针同时指向两个字符串的头部
2. 开始移动，如果当前s的索引字节和t 当前索引字节相同 s索引向后移动
3. 而t不管找没有找到需要移动