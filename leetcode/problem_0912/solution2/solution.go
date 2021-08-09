package solution

// insert sort
func sortArray(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		e := nums[i]
		var j int
		for j = i; j > 0 && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}
	return nums
}
