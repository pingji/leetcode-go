package solution

func isVaild(row int, col int, n int, path [][]string) bool {
	i := 0
	for ; i < row; i++ {
		for j := 0; j < n; j++ {
			// top left \
			if j == col-(row-i) && path[i][j] == "Q" {
				return false
			}
			// top right /
			if j == col+(row-i) && path[i][j] == "Q" {
				return false
			}
			// column
			if j == col && path[i][j] == "Q" {
				return false
			}
		}
	}
	// i == row
	for j := 0; j <= col; j++ {
		// row
		if path[i][j] == "Q" {
			return false
		}
	}
	return true
}

func totalNQueens(n int) int {
	var res = 0
	path := make([][]string, n)
	for i := range path {
		path[i] = make([]string, n)
		for j := range path[i] {
			path[i][j] = "."
		}
	}
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}
		for col := 0; col < n; col++ {
			if !isVaild(row, col, n, path) {
				continue
			}
			path[row][col] = "Q"

			backtrack(row + 1)

			path[row][col] = "."
		}
	}
	backtrack(0)

	return res
}
