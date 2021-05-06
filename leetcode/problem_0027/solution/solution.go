package solution

func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[j] = nums[i]
			}
			j++
		}
	}
	return j
}
