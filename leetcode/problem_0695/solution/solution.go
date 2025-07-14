package solution

// maxAreaOfIsland 计算岛屿的最大面积
// 使用深度优先搜索(DFS)算法，时间复杂度 O(m*n)，空间复杂度 O(m*n)
// 算法思想：遍历网格，对每个岛屿计算其面积，记录最大面积
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0 // 记录最大岛屿面积
	// 遍历整个网格
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 遇到陆地时，计算该岛屿的面积
			if grid[i][j] == 1 {
				area := dfs(grid, i, j)      // DFS计算岛屿面积
				maxArea = max(area, maxArea) // 更新最大面积
			}
		}
	}

	return maxArea
}

// dfs 深度优先搜索，计算从位置(x,y)开始的岛屿面积
// 参数：grid - 网格，x, y - 当前访问的位置
// 返回值：该岛屿的面积
func dfs(grid [][]int, x int, y int) int {
	// 边界条件：超出网格范围或遇到水
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	// 将当前陆地标记为水，避免重复访问
	grid[x][y] = 0
	// 当前陆地格子贡献1个面积单位
	area := 1 
	// 向四个方向递归计算面积：上、左、下、右
	area += dfs(grid, x-1, y) // 向上
	area += dfs(grid, x, y-1) // 向左
	area += dfs(grid, x+1, y) // 向下
	area += dfs(grid, x, y+1) // 向右
	return area
}

// max 返回两个整数中的较大值
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
