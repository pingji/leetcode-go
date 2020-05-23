package problem0053

// solution-1: DP
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	maxsum := sum[0]

	for i := 1; i < len(nums); i++ {
		if sum[i-1] > 0 {
			sum[i] = sum[i-1] + nums[i]
		} else {
			sum[i] = nums[i]
		}
		if sum[i] > maxsum {
			maxsum = sum[i]
		}
	}
	return maxsum
}
