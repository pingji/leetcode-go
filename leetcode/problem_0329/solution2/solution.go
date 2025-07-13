package solution

// 四个方向的移动：上、左、右、下
var (
	dirs          = [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	rows, columns int
)

// longestIncreasingPath 计算矩阵中最长递增路径的长度
// 使用记忆化DFS，避免重复计算
func longestIncreasingPath(matrix [][]int) int {
	// 边界检查
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])

	// memo[i][j] 表示从点(i,j)开始的最长递增路径长度
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}

	// 从每个点开始DFS，找到最长的递增路径
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, memo))
		}
	}

	return ans
}

// dfs 从点(row, column)开始DFS，返回最长递增路径长度
// 使用记忆化避免重复计算
func dfs(matrix [][]int, row, column int, memo [][]int) int {
	// 如果已经计算过，直接返回
	if memo[row][column] > 0 {
		return memo[row][column]
	}

	// 初始长度为1（包含当前点）
	memo[row][column]++

	// 尝试四个方向
	for _, dir := range dirs {
		newRow, newColumn := row+dir[0], column+dir[1]
		// 检查邻居是否在边界内且值更大
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
			// 更新最长路径长度
			memo[row][column] = max(memo[row][column], dfs(matrix, newRow, newColumn, memo)+1)
		}
	}
	return memo[row][column]
}

// max 返回两个整数中的较大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
