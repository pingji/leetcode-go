# Solutions

## Solution 1：深度优先遍历

### 算法思路

使用深度优先搜索(DFS)来计算岛屿的周长：

1. **遍历网格**：遍历整个二维网格的每个位置
2. **发现陆地**：当遇到值为 1 的位置时，说明发现了一个陆地格子
3. **DFS计算周长**：从该位置开始进行深度优先搜索，计算该陆地格子贡献的周长
4. **累加周长**：将所有陆地格子贡献的周长累加得到总周长

### 核心算法

```go
func islandPerimeter(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    rows, cols := len(grid), len(grid[0])
    perimeter := 0
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                perimeter += dfs(grid, i, j)
            }
        }
    }
    
    return perimeter
}
```

### DFS函数

```go
func dfs(grid [][]int, i, j int) int {
    // 边界条件：超出网格范围，贡献1条边
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return 1
    }
    
    // 遇到水，贡献1条边
    if grid[i][j] == 0 {
        return 1
    }
    
    // 已经访问过的陆地，不贡献周长
    if grid[i][j] == 2 {
        return 0
    }
    
    // 标记当前陆地为已访问
    grid[i][j] = 2
    
    // 向四个方向递归计算周长
    return dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
}
```

### 周长计算原理

每个陆地格子对周长的贡献取决于其与水域或边界的接触面：

1. **边界接触**：陆地格子与网格边界接触，贡献1条边
2. **水域接触**：陆地格子与水域(值为0)接触，贡献1条边
3. **陆地接触**：陆地格子与相邻陆地接触，不贡献周长
4. **重复访问**：已经访问过的陆地格子不重复计算

### 关键要点

1. **标记机制**：使用值2标记已访问的陆地，避免重复计算
2. **边界处理**：超出网格范围时返回1，表示贡献1条边
3. **递归计算**：向四个方向递归，累加所有贡献的周长
4. **避免重复**：已访问的陆地格子直接返回0

### 复杂度分析

- **时间复杂度**：O(m×n)，其中 m 和 n 分别是网格的行数和列数
- **空间复杂度**：O(m×n)，最坏情况下整个网格都是陆地，递归深度为 O(m×n)

# References

- [岛屿类问题的通用解法、DFS 遍历框架](https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-)

- [「手画图解」463. 岛屿的周长 | 很简单的解法](https://leetcode.cn/problems/island-perimeter/solutions/466367/shou-hua-tu-jie-463-dao-yu-de-zhou-chang-by-xiao_b)