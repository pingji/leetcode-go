## Solutions
### solution 1
回溯

对比问题78，包含重复元素的数组的子集，只不过一个小的改进。

比如说求 nums = [1,2,2] 的子集，那么对于子集 [1,2] 是选择了第一个 2，那么就不能再选第二个 2 来构成 [1,2] 了。所以，此时的改动点，就是先排序，每个元素 nums[i] 添加到 path 之前，判断一下 nums[i] 是否等于 nums[i - 1] ，如果相等就不添加到 path 中。

本题也可以不使用used数组来去重，因为递归的时候下一个startIndex是i+1而不是0。如果要是全排列的话，每次要从0开始遍历，为了跳过已入栈的元素，需要使用used。

```
// 第一层选择 (begin = 0)
i = 0: nums[0] = 1, 选择 1
i = 1: nums[1] = 2, 选择 2  
i = 2: nums[2] = 2, 检查 nums[2] == nums[1] (2 == 2) 且 i > begin (2 > 0)
        → 跳过，避免重复选择第二个2

// 第二层选择 (begin = 1)
i = 1: nums[1] = 2, 选择 2
i = 2: nums[2] = 2, 检查 nums[2] == nums[1] (2 == 2) 且 i > begin (2 > 1)
        → 跳过，避免重复选择第二个2
```

# References

- [代码随想录」带你学透回溯算法！90. 子集 II:【彻底理解子集问题如何去重】](https://leetcode.cn/problems/subsets-ii/solutions/690866/90-zi-ji-iiche-di-li-jie-zi-ji-wen-ti-ru-djmf)