package solution

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return size
	}
	i, j := 2, 2
	for i < len(nums) {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}
