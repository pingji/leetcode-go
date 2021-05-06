package solution

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(nums)

	var backtrack func(int)
	backtrack = func(begin int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := begin; i < len(nums); i++ {
			if i > begin && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)

	return res
}
