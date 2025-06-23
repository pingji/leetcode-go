package solution2

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxarea := 0

	for i < j {
		// 计算当前面积
		area := min(height[i], height[j]) * (j - i)
		maxarea = max(maxarea, area)

		// 关键优化：跳过不可能产生更大面积的指针移动
		if height[i] < height[j] {
			// 移动短板，但跳过所有比当前短板更短的柱子
			currHeight := height[i]
			for i < j && height[i] <= currHeight {
				i++
			}
		} else {
			// 移动短板，但跳过所有比当前短板更短的柱子
			currHeight := height[j]
			for i < j && height[j] <= currHeight {
				j--
			}
		}
	}

	return maxarea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
