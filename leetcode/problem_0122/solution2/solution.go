package solution

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		if delta > 0 {
			profit += delta
		}
	}
	return profit
}
