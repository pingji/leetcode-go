package solution

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
