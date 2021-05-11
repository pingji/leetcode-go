package solution

func sortColors(nums []int) {
	// [0, p0) = 0
	// [p0, i) = 1
	// (p2, len-1] = 2

	p0 := 0
	p2 := len(nums) - 1
	i := 0
	for i <= p2 {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		}
	}
}
