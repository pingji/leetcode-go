package solution

var (
	dps []int
)

func coinChange(coins []int, amount int) int {
	dps = make([]int, amount+1)
	return dp(amount, coins)

}

func dp(i int, coins []int) int {
	if i < 0 {
		return -1
	}
	if i == 0 {
		return 0
	}
	if dps[i] != 0 {
		return dps[i]
	}
	min := i + 1
	for _, v := range coins {
		x := dp(i-v, coins)
		if x >= 0 && x < min {
			min = dp(i-v, coins)
		}
	}
	if min == i+1 {
		dps[i] = -1
	} else {
		dps[i] = min + 1
	}
	return dps[i]
}
