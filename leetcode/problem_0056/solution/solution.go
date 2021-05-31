package solution

import "sort"

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 {
		return res
	}

	// sort
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= cur[1] {
			r := max(cur[1], intervals[i][1])
			cur[1] = r
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)

	return res
}
