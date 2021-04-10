package solution

func mySqrt(x int) int {
	l, r := 0, x
	ret := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
