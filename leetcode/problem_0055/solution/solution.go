package solution

func canJump(nums []int) bool {
	maxIndex := 0
	for i, v := range nums {
		if i <= maxIndex {
			maxIndex = max(maxIndex, i+v)
		} else {
			break
		}
	}
	if maxIndex >= len(nums)-1 {
		return true
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
