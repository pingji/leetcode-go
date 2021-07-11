# Solutions

## Solution 1：贪心算法
整体最优：整个序列有最多的局部峰值，从而达到最长摆动序列

实际操作上，其实连删除的操作都不用做，因为题目要求的是最长摆动子序列的长度，所以只需要统计数组的峰值数量就可以了（相当于是删除单一坡度上的节点，然后统计长度）

# References
- [贪心算法：摆动序列](https://mp.weixin.qq.com/s?__biz=MzUxNjY5NTYxNA==&mid=2247485801&idx=1&sn=b9d69b9df171995701540c18d671a12b&chksm=f9a23a38ced5b32e83d5300ee43b668e397047f201380af65bd0c4d65fb39684cf7971d3d0e2&cur_album_id=1485825793120387074&scene=189#rd)