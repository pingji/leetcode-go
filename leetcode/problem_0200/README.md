# Solutions

## Solution 1：深度优先遍历

### 算法思路

使用深度优先搜索(DFS)来解决岛屿数量问题：

1. **遍历网格**：遍历整个二维网格的每个位置
2. **发现陆地**：当遇到值为 '1' 的位置时，说明发现了一个岛屿的起点
3. **DFS标记**：从该位置开始进行深度优先搜索，标记整个岛屿
4. **计数统计**：每完成一次DFS，岛屿数量加1

### 核心算法

```go
func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)  // DFS标记整个岛屿
                count++          // 岛屿数量加1
            }
        }
    }
    return count
}
```

### DFS函数

```go
func dfs(grid [][]byte, x int, y int) {
    // 边界条件：超出网格范围或遇到水
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
        return
    }
    // 将当前陆地标记为水，避免重复访问
    grid[x][y] = '0'
    // 向四个方向进行DFS：上、左、下、右
    dfs(grid, x-1, y) // 向上
    dfs(grid, x, y-1) // 向左
    dfs(grid, x+1, y) // 向下
    dfs(grid, x, y+1) // 向右
}
```

### 关键要点

1. **标记访问**：将访问过的陆地标记为水('0')，避免重复访问
2. **边界检查**：DFS时需要检查是否超出网格边界
3. **四个方向**：向上下左右四个方向进行搜索
4. **岛屿定义**：相邻的陆地(上下左右)属于同一个岛屿

### 复杂度分析

- **时间复杂度**：O(m×n)，其中 m 和 n 分别是网格的行数和列数
- **空间复杂度**：O(m×n)，最坏情况下整个网格都是陆地，递归深度为 O(m×n)

## Solution 2：并查集

# References

- [200. 岛屿数量（DFS / BFS）](https://leetcode-cn.com/problems/number-of-islands/solution/number-of-islands-shen-du-you-xian-bian-li-dfs-or-/)
- [岛屿类问题的通用解法、DFS 遍历框架](https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-)
