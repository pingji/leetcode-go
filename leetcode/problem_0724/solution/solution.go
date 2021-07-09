package solution

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i, num := range nums {
		if (leftSum*2 + num) == sum {
			return i
		}
		leftSum += num
	}
	return -1
}
