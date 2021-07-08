# Solutions

## Solution 1：前缀和 + 哈希表优化
我们建立哈希表 \textit{mp}mp，以和为键，出现次数为对应的值，记录 \textit{pre}[i]pre[i] 出现的次数，从左往右边更新 \textit{mp}mp 边计算答案，那么以 ii 结尾的答案 \textit{mp}[\textit{pre}[i]-k]mp[pre[i]−k] 即可在 O(1)O(1) 时间内得到。

参考
- [前缀和的运用，一步步优化 | 560. 和为K的子数组](https://leetcode-cn.com/problems/subarray-sum-equals-k/solution/dai-ni-da-tong-qian-zhui-he-cong-zui-ben-fang-fa-y/)
- [秒懂 前缀和+hashmap](https://leetcode-cn.com/problems/subarray-sum-equals-k/solution/xue-sheng-wu-de-nu-peng-you-du-neng-kan-5fn1m/)