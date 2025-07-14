package solution

// islandPerimeter 计算岛屿的周长
// 使用深度优先搜索(DFS)算法，时间复杂度 O(m*n)，空间复杂度 O(m*n)
// 算法思想：遍历网格，对每个陆地格子计算其贡献的周长
func islandPerimeter(grid [][]int) int {
	// 边界检查：如果网格为空，返回0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	perimeter := 0 // 周长计数器

	// 遍历整个网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 遇到陆地时，计算该陆地格子贡献的周长
			if grid[i][j] == 1 {
				perimeter += dfs(grid, i, j)
			}
		}
	}

	return perimeter
}

// dfs 深度优先搜索，计算从位置(i,j)开始的陆地格子贡献的周长
// 参数：grid - 网格，i, j - 当前访问的位置
// 返回值：该位置贡献的周长
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

	// 向四个方向递归计算周长：上、下、左、右
	return dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
}
