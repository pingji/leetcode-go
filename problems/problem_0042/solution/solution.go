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
	sum := 0
	size := len(height)
	for i := 0; i < size; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j < size; j++ {
			maxRight = max(maxRight, height[j])
		}
		sum += min(maxLeft, maxRight) - height[i]
	}
	return sum
}
