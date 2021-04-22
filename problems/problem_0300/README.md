# Solutions
## Solution 1: 动态规划
定义 $df[i]$  为考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度，注意 $nums[i]$ 必须被选取。

$df[i]=\max(df[j])+1,df[i])$ ,其中 $0≤j<i$ 且 $nums[j]<nums[i]$

复杂度
- 时间复杂度：$O(n^2)$
- 空间复杂度：$O(n)$