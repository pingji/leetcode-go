package leetcode

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
    dp := make([]int, n+1)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i <= n; i++ {
		if i != n {
			dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
		} else {
			dp[i] = min(dp[i-2], dp[i-1]) + 0
		}
	}
	return dp[n]
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}