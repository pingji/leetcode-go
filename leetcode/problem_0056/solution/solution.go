package solution

import "sort"

// max 返回两个整数中的较大值
// 参数：x, y - 两个整数
// 返回值：较大的整数
func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

// merge 合并重叠的区间
// 参数：intervals - 区间数组，每个区间用 [start, end] 表示
// 返回值：合并后的区间数组
// 算法思路：
// 1. 按区间起始位置排序
// 2. 遍历排序后的区间，合并重叠的区间
// 时间复杂度：O(nlogn)，主要来自排序
// 空间复杂度：O(n)，存储结果数组
func merge(intervals [][]int) [][]int {
	res := [][]int{}
	// 处理空数组的情况
	if len(intervals) == 0 {
		return res
	}

	// 按区间起始位置进行排序
	// 这样可以保证相邻的区间更容易合并
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化当前区间为第一个区间
	cur := intervals[0]
	// 遍历剩余的区间
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的起始位置小于等于前一个区间的结束位置，说明有重叠
		if intervals[i][0] <= cur[1] {
			// 合并区间：取两个区间结束位置的最大值作为新的结束位置
			r := max(cur[1], intervals[i][1])
			cur[1] = r
		} else {
			// 没有重叠，将当前区间加入结果，并更新当前区间
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	// 将最后一个区间加入结果
	res = append(res, cur)

	return res
}
