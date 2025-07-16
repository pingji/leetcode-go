package solution

// isVaild 检查在指定位置放置皇后是否有效
// row: 当前行
// col: 当前列
// n: 棋盘大小
// path: 当前棋盘状态
func isVaild(row int, col int, n int, path [][]string) bool {
	i := 0
	// 检查上方所有行是否有冲突
	for ; i < row; i++ {
		for j := 0; j < n; j++ {
			// 检查左上对角线 \ (row-i == col-j)
			if row-i == col-j && path[i][j] == "Q" {
				return false
			}
			// 检查右上对角线 / (row-i == j-col)
			if row-i == j-col && path[i][j] == "Q" {
				return false
			}
			// 检查同一列
			if j == col && path[i][j] == "Q" {
				return false
			}
		}
	}
	// 检查当前行的左侧是否有皇后
	for j := 0; j <= col; j++ {
		// 检查同一行
		if path[i][j] == "Q" {
			return false
		}
	}
	return true
}

// totalNQueens 计算N皇后问题的解的数量
// 与N皇后I的区别：只返回解的数量，不需要返回具体的解
func totalNQueens(n int) int {
	var res = 0 // 记录解的数量

	// 初始化棋盘，所有位置都设为"."
	path := make([][]string, n)
	for i := range path {
		path[i] = make([]string, n)
		for j := range path[i] {
			path[i][j] = "."
		}
	}

	// 回溯函数，逐行放置皇后
	var backtrack func(int)
	backtrack = func(row int) {
		// 如果已经放置了n个皇后，说明找到一个解
		if row == n {
			res++ // 解的数量加1
			return
		}

		// 在当前行的每一列尝试放置皇后
		for col := 0; col < n; col++ {
			// 检查当前位置是否可以放置皇后
			if !isVaild(row, col, n, path) {
				continue
			}

			// 放置皇后
			path[row][col] = "Q"

			// 递归处理下一行
			backtrack(row + 1)

			// 回溯：移除皇后，尝试下一个位置
			path[row][col] = "."
		}
	}

	// 从第0行开始回溯
	backtrack(0)

	return res
}
