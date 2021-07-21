package solution

func numWays(n int) int {
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		a, b = b%1000000007, (a+b)%1000000007
	}
	return a
}
