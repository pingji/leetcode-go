package solution

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, i, j)
				maxArea = max(area, maxArea)
			}
		}
	}

	return maxArea
}

func dfs(grid [][]int, x int, y int) int {
	area := 0
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return area
	}
	grid[x][y] = 0
	area += 1
	area += dfs(grid, x-1, y)
	area += dfs(grid, x, y-1)
	area += dfs(grid, x+1, y)
	area += dfs(grid, x, y+1)
	return area
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
