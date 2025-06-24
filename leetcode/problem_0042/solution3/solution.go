package solution

// trap 使用单调栈解决接雨水问题
// 单调栈维护递减的高度索引，当遇到更高的柱子时计算积水
func trap(height []int) int {
	sum := 0
	// 单调栈，存储递减高度的索引
	stack := make([]int, 0)

	for i := 0; i < len(height); i++ {
		// 当栈不为空且当前高度大于栈顶高度时，说明可以形成积水
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 弹出栈顶元素，作为积水的底部
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 如果栈中还有元素，说明左边有更高的柱子，可以形成积水
			if len(stack) > 0 {
				// 计算积水高度：左右两边柱子的较小值减去底部高度
				// 计算积水宽度：当前位置减去左边柱子的位置再减1
				sum += (min(height[i], height[stack[len(stack)-1]]) - height[top]) * (i - stack[len(stack)-1] - 1)
			}
		}
		// 将当前索引压入栈中
		stack = append(stack, i)
	}
	return sum
}

// min 返回两个整数中的较小值
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
