package solution

func generateParenthesis(n int) []string {
	res := []string{}
	var backtrack func(path string, left int, right int)
	backtrack = func(path string, left int, right int) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		if left < n {
			backtrack(path+"(", left+1, right)
		}
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}
	backtrack("", 0, 0)

	return res
}
