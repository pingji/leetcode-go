package solution

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var res float64 = 1
	for i := 0; i < n; i++ {
		res = res * x
	}
	return res
}
