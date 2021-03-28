package solution

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minprice := prices[0]
	maxprofit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
		if prices[i] < minprice {
			minprice = prices[i]
		}
	}

	return maxprofit
}
