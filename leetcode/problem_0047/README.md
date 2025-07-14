## Solutions

### solution1

采用回溯算法,回溯算法的框架为

```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```

需要注意的是，result.add(路径) 需要深度拷贝

```mermaid
graph TB
    root[1,1,2 ] --> Node11[1]
    root -.-> Node12[1]
    root --> Node13[2]
    Node11 --> Node21[1]
    Node11 --> Node22[2]
    Node12 -.-> Node23[1]
    Node12 -.-> Node24[2]
    Node13 --> Node25[1]
    Node13 -.-> Node26[1]
    Node21 --> Node31[2]
    Node22 --> Node32[1]
    Node23 -.-> Node33[2]
    Node24 -.-> Node34[1]
    Node25 --> Node35[1]
    Node26 -.-> Node36[1]
```

主要讲一下为什么剪枝条件加了一个!used[i-1]？其实图里也很清楚，我们知道重复必定发生在决策树的同一层，当我们只有i>0 && nums[i] == nums[i-1]时，有一种可能是nums[i-1]和nums[i]在同一条路径上，nums[i-1]是num[i]的父节点，对应图中左下的箭头，是生成[1,1*]的那条路，可以重复选。而我们需要剪的是回溯到选第一个位置时，选到第二个1，这个和之前的那个1是同一层的，而且是重复的，因此必须再加上一个!used[i-1]来保证不是重复选择[1,1*]的情况确实是同一层选重复了。

假设我们有 3 个重复数排完序后相邻，那么我们一定保证每次都是拿从左往右第一个未被填过的数字，即整个数组的状态其实是保证了 [未填入，未填入，未填入] 到 [填入，未填入，未填入]，再到 [填入，填入，未填入]，最后到 [填入，填入，填入] 的过程的，因此可以达到去重的目标。

# References

- [47.排列问题回溯，剪枝重复数的思考](https://leetcode-cn.com/problems/permutations-ii/solution/47hui-su-jian-zhi-zhong-fu-shu-by-bu-zhi-h7dm/)
- [全排列 II](https://leetcode.cn/problems/permutations-ii/solutions/417937/quan-pai-lie-ii-by-leetcode-solution)
