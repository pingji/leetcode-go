package solution

import "sort"

// permuteUnique 返回数组的所有不重复全排列
// 使用回溯算法，通过排序和剪枝避免重复排列
func permuteUnique(nums []int) [][]int {
	// res 存储所有不重复的排列结果
	res := [][]int{}
	// path 存储当前正在构建的排列
	path := []int{}
	// 先排序，便于后续去重
	sort.Ints(nums)
	// flags 标记数组，记录每个数字是否已被使用
	flags := make([]bool, len(nums))

	// backtrack 回溯函数，生成所有不重复的排列
	var backtrack func()
	backtrack = func() {
		// 如果当前路径长度等于原数组长度，说明找到了一个完整排列
		if len(nums) == len(path) {
			// 创建path的副本，避免后续修改影响已保存的结果
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		// 遍历所有数字，选择未使用的数字
		for index, num := range nums {
			// 如果当前数字已被使用，跳过
			if flags[index] {
				continue
			}
			// 去重剪枝：如果当前数字与前一个数字相同，且前一个数字未被使用，则跳过
			// 这确保了相同数字的相对顺序，避免重复排列
			if index > 0 && nums[index] == nums[index-1] && !flags[index-1] {
				continue
			}

			// 选择当前数字
			path = append(path, num)
			// 标记当前数字为已使用
			flags[index] = true

			// 递归生成剩余数字的排列
			backtrack()

			// 回溯：撤销选择，恢复状态
			path = path[:len(path)-1]
			flags[index] = false
		}
	}

	// 开始回溯搜索
	backtrack()

	return res
}
