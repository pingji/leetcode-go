package solution

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	// 0: none
	// 1: buy_1
	// 2: buy_1 sell_1
	// 3: buy_1 sell_1 buy_2
	// 4: buy_1 sell_1 buy_2 sell_2
	dp := make([][5]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[n-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
