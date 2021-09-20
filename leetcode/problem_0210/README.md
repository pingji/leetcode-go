# Solutions

## Solution 1：拓扑排序
需要以下几个数据结构
1. edges := make([][]int, numCourses)  记录某个点有多少个依赖它的点
2. indeg := make([]int, numCourses) 记录每个点的入度
3. result := []int{}  记录结果
4. q := []int{}   先入先出的队列

判断 if len(result) == numCourses 是否全部都进入result

# References
