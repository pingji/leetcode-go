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
只search一遍，用 双指针法即可。如果需要search多次，就需要对t中各个字母的位置建立索引。

建立索引有多种方法
1. 建立一个矩阵，参考solution 2 动态规划
2. 建立一个哈希表，参考solution 3 二分查找

### Solution 1:  Two Pointers
解题思路
1. 两个指针同时指向两个字符串的头部
2. 开始移动，如果当前s的索引字节和t 当前索引字节相同 s索引向后移动
3. 而t不管找没有找到需要移动

复杂度
- 时间复杂度：O(m+n)。m,n分别为s和t的长度
- 空间复杂度：O(1)

### Solution 2：Dynamic Programming
令 f[i][j] 表示字符串 t 中从位置 i 开始往后字符 j 第一次出现的位置。
转移方程：
$$ 
f[i][j] =
\begin{cases}
i \qquad & t[i] = j \\
f[i+1][j] \qquad & t[i] \neq j
\end{cases}
$$


### Solution 3：Binary Search

