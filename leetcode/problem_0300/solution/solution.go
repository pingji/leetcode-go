package solution

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	maxres := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		maxres = max(maxres, dp[i])
	}

	return maxres
}
