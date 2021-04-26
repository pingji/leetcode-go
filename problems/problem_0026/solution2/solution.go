package solution

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	i, j := 1, 1
	for i < len(nums) {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}
