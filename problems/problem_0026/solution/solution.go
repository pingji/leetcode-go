package solution

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
