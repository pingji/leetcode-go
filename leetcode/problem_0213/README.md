# Solutions

## Solution 1：Dynamic Programming

```
dp[i] = max(dp[i−2]+nums[i],dp[i−1])
```
分别取 (start, end}=(0,n-2), 和 (start,end)=(1,n−1) 进行计算，取两个 dp[end] 中的最大值，即可得到最终结果。
