package solution

// 四个方向的移动：上、左、右、下
var (
	dirs = [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
)

// longestIncreasingPath 计算矩阵中最长递增路径的长度
// 使用拓扑排序的思想，从出度为0的点开始，逐层计算
func longestIncreasingPath(matrix [][]int) int {
	// 边界检查
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns := len(matrix), len(matrix[0])

	// outdegrees[i][j] 表示从点(i,j)出发可以到达的更大值的点的数量
	outdegrees := make([][]int, rows)
	for i := 0; i < rows; i++ {
		outdegrees[i] = make([]int, columns)
	}

	// 计算每个点的出度（可以到达的更大值的邻居数量）
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for _, dir := range dirs {
				newRow, newColumn := i+dir[0], j+dir[1]
				// 检查邻居是否在边界内且值更大
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[i][j] {
					outdegrees[i][j]++ // 出度+1
				}
			}
		}
	}

	// 将所有出度为0的点加入队列（这些点是递增路径的终点）
	queue := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if outdegrees[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	// 逐层处理，ans记录层数（即最长递增路径的长度）
	ans := 0
	for len(queue) != 0 {
		ans++ // 每处理一层，路径长度+1
		size := len(queue)
		// 处理当前层的所有点
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]
			row, column := cell[0], cell[1]

			// 检查所有邻居，如果邻居值更小，则减少邻居的出度
			for _, dir := range dirs {
				newRow, newColumn := row+dir[0], column+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] < matrix[row][column] {
					outdegrees[newRow][newColumn]--         // 邻居出度-1
					if outdegrees[newRow][newColumn] == 0 { // 如果邻居出度变为0，加入下一层
						queue = append(queue, []int{newRow, newColumn})
					}
				}
			}
		}
	}
	return ans
}
