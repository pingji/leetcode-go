package solution

// combine 返回从1到n中选择k个数的所有组合
// 使用回溯算法，通过递归生成所有可能的组合
func combine(n int, k int) [][]int {
	// res 存储所有组合结果
	res := [][]int{}
	// path 存储当前正在构建的组合
	path := []int{}

	// backtrack 回溯函数，begin表示当前可以选择的起始数字
	var backtrack func(int)
	backtrack = func(begin int) {
		// 如果当前路径长度等于k，说明找到了一个有效组合
		if len(path) == k {
			// 创建path的副本，避免后续修改影响已保存的结果
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// 从begin开始选择数字，避免重复组合
		for i := begin; i <= n; i++ {
			// 选择当前数字
			path = append(path, i)
			// 递归选择下一个数字，从i+1开始避免重复
			backtrack(i + 1)
			// 回溯：撤销选择，尝试其他数字
			path = path[:len(path)-1]
		}
	}

	// 从数字1开始回溯搜索
	backtrack(1)

	return res
}
