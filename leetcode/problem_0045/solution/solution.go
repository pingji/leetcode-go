package solution

func jump(nums []int) int {
	maxIndex := 0
	steps := 0
	curMaxIndex := 0
	curIndex := 0
	for i, v := range nums {
		maxIndex = max(maxIndex, i+v)
		if i == curMaxIndex {
			curMaxIndex = maxIndex
			curIndex = i
			if i > 0 {
				steps++
			}
		}
	}
	if curIndex < len(nums)-1 {
		steps++
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
