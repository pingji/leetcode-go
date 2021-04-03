package solution

func permute(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	flags := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(nums) == len(path) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for index, num := range nums {
			if !flags[index] {
				path = append(path, num)
				flags[index] = true

				backtrack()

				path = path[:len(path)-1]
				flags[index] = false
			}
		}
	}

	backtrack()

	return res
}
