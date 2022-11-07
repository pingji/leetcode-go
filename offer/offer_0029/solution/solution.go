package solution

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	res := make([]int, rows*columns)
	index := 0
	left, right, top, bottom := 0, columns-1, 0, rows-1

	for {
		// top: left -> right
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		top++
		if top > bottom {
			break
		}
		// right: top ->bottom
		for i := top; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		right--
		if left > right {
			break
		}
		// bottom: right->left
		for i := right; i >= left; i-- {
			res[index] = matrix[bottom][i]
			index++
		}
		bottom--
		if top > bottom {
			break
		}
		// left: bottom->top
		for i := bottom; i >= top; i-- {
			res[index] = matrix[i][left]
			index++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
