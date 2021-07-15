package solution

func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, row int, col int) bool {
	if col == len(board[0]) {
		return backtrack(board, row+1, 0)
	}
	if row == len(board) {
		return true
	}
	if board[row][col] != '.' {
		return backtrack(board, row, col+1)
	}
	for k := byte('1'); k <= byte('9'); k++ {
		if isValid(board, row, col, k) {
			board[row][col] = k
			if backtrack(board, row, col+1) {
				return true
			} else {
				board[row][col] = '.'
			}
		}
	}
	return board[row][col] != '.'
}

func isValid(board [][]byte, row int, col int, val byte) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == val {
			return false
		}
	}
	for j := 0; j < len(board[row]); j++ {
		if board[row][j] == val {
			return false
		}
	}
	beginRow := (row / 3) * 3
	beginCol := (col / 3) * 3

	for i := beginRow; i < beginRow+3; i++ {
		for j := beginCol; j < beginCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
