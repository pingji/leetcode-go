package solution

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gi := len(g) - 1
	si := len(s) - 1
	res := 0

	for gi >= 0 && si >= 0 {
		if s[si] >= g[gi] {
			gi--
			si--
			res++
		} else {
			gi--
		}
	}

	return res
}
