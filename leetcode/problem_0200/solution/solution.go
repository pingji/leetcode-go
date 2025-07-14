package solution

// numIslands 计算岛屿数量
// 使用深度优先搜索(DFS)算法，时间复杂度 O(m*n)，空间复杂度 O(m*n)
// 算法思想：遍历网格，遇到陆地时进行DFS标记整个岛屿，统计岛屿数量
func numIslands(grid [][]byte) int {
	count := 0 // 岛屿计数器
	// 遍历整个网格
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 遇到陆地时，进行DFS标记整个岛屿
			if grid[i][j] == '1' {
				dfs(grid, i, j) // DFS标记整个岛屿
				count++         // 岛屿数量加1
			}
		}
	}

	return count
}

// dfs 深度优先搜索，标记整个岛屿
// 参数：grid - 网格，x, y - 当前访问的位置
// 算法：将访问过的陆地标记为水，避免重复访问
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
