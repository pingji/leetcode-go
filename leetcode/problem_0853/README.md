# Solutions
## Solution 1

### 解题思路

这是一道关于车队形成的题目。核心思路是：**从后往前遍历车辆，如果后面的车到达目标点的时间小于等于前面的车，则会被前面的车阻挡，形成车队**。

#### 算法步骤：

1. **数据结构设计**：
   - 定义 `Car` 结构体，包含位置 `pos` 和到达目标点的时间 `time`

2. **计算到达时间**：
   - 对每辆车计算到达目标点的时间：`time = (target - position) / speed`
   - 这个时间表示车辆在不受阻挡的情况下到达目标点所需的时间

3. **排序处理**：
   - 按位置从大到小排序（从后往前）
   - 这样可以从最接近目标的车辆开始处理

4. **车队统计**：
   - 从第二辆车开始遍历
   - 如果当前车到达时间 ≤ 前一辆车到达时间：当前车会被阻挡，时间更新为前一辆车的时间
   - 如果当前车到达时间 > 前一辆车到达时间：当前车不会被阻挡，形成新的车队，计数+1

#### 时间复杂度：O(n log n)，空间复杂度：O(n)

### 参考
- [go 解题](https://leetcode-cn.com/problems/car-fleet/solution/go-yu-qi-ta-lao-ge-si-ji-ben-lu-yi-zhi-by-yu-hu-yi/)
- [官方解题](https://leetcode-cn.com/problems/car-fleet/solution/che-dui-by-leetcode/)