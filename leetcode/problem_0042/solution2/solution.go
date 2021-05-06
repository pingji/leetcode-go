package solution

func max(a int, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func trap(height []int) int {
	size := len(height)
	if size < 1 {
		return 0
	}

	maxLeftArr := make([]int, size)
	maxLeftArr[0] = height[0]
	for i := 1; i < size; i++ {
		maxLeftArr[i] = max(maxLeftArr[i-1], height[i])
	}

	maxRighttArr := make([]int, size)
	maxRighttArr[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRighttArr[i] = max(maxRighttArr[i+1], height[i])
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum += min(maxLeftArr[i], maxRighttArr[i]) - height[i]
	}

	return sum
}
