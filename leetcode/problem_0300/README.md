# Solutions
## Solution 1: 动态规划
定义 $df[i]$  为考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度，注意 $nums[i]$ 必须被选取。

$df[i]=\max(df[j])+1,df[i])$ ,其中 $0≤j<i$ 且 $nums[j]<nums[i]$

复杂度
- 时间复杂度：$O(n^2)$
- 空间复杂度：$O(n)$

## Solution 2: 贪心算法 + 二分查找
参考 [动态规划 - 包含O(N log N) 解法的状态定义以及解释](https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/)

$tail[i]$ 表示：长度为 i + 1 的 所有 上升子序列的结尾的最小值。 数组 tail 也是一个严格上升数组，通过二分查找降低时间复杂度

在有序数组 tail 中查找第 1 个等于大于 num 的那个数，试图让它变小；
- 如果有序数组 tail 中存在 等于 num 的元素，什么都不做，因为以 num 结尾的最短的「上升子序列」已经存在；
- 如果有序数组 tail 中存在 大于 num 的元素，找到第 1 个，让它变小，这样我们就找到了一个 结尾更小的相同长度的上升子序列。


复杂度
- 时间复杂度：$O(n\log{n})$, 遍历数组使用了 $O(n)$，二分查找法使用了  $O(\log{n})$
- 空间复杂度：$O(n)$
