package solution

func dailyTemperatures(temperatures []int) []int {
	stack := []int{} // 存储索引
	ans := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		// 当栈不为空且当前温度大于栈顶温度时，更新答案
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 栈顶元素的等待天数 = 当前索引 - 栈顶索引
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
		// 将当前索引入栈
		stack = append(stack, i)
	}

	return ans
}
