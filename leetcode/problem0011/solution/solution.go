package solution

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxarea := 0
	area := 0
	for i < j {
		if height[i] > height[j] {
			area = height[j] * (j - i)
			j--
		} else {
			area = height[i] * (j - i)
			i++
		}
		if area > maxarea {
			maxarea = area
		}
	}
	return maxarea
}
