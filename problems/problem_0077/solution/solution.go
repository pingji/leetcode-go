package solution

func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(begin int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := begin; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)

	return res
}
