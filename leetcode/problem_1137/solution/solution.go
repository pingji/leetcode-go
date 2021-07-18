package solution

func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	switch n {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	default:
		for i := 3; i <= n; i++ {
			a, b, c = b, c, a+b+c
		}
		return c
	}
}
