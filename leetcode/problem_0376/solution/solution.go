package solution

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if len(nums) < 2 {
		return n
	}
	ans := 1
	preDiff := 0
	curDiff := 0
	for i := 1; i < n; i++ {
		curDiff = nums[i] - nums[i-1]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			ans++
			preDiff = curDiff
		}
	}

	return ans
}
