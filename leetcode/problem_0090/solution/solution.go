package solution

import "sort"

// subsetsWithDup 返回数组的所有不重复子集
// 使用回溯算法，通过排序和剪枝避免重复子集
func subsetsWithDup(nums []int) [][]int {
	// res 存储所有不重复的子集结果
	res := [][]int{}
	// path 存储当前正在构建的子集
	path := []int{}
	// 先排序，便于后续去重
	sort.Ints(nums)

	// backtrack 回溯函数，begin表示当前可以选择的起始位置
	var backtrack func(int)
	backtrack = func(begin int) {
		// 每个节点都保存当前路径作为子集（包括空集）
		// 这确保了所有可能的子集都被包含
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		// 从begin开始选择数字，避免重复子集
		for i := begin; i < len(nums); i++ {
			// 去重剪枝：如果当前数字与前一个数字相同，且不是第一次选择，则跳过
			// 这确保了相同数字的相对顺序，避免重复子集
			if i > begin && nums[i] == nums[i-1] {
				continue
			}

			// 选择当前数字
			path = append(path, nums[i])
			// 递归选择下一个数字，从i+1开始避免重复
			backtrack(i + 1)
			// 回溯：撤销选择，尝试其他数字
			path = path[:len(path)-1]
		}
	}

	// 从索引0开始回溯搜索
	backtrack(0)

	return res
}
