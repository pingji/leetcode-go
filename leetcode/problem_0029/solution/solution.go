package solution

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if divisor == 0 {
		return 0
	}

	negative := false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		negative = true
	}
	divd, divs := dividend, divisor
	if dividend < 0 {
		divd = -dividend
	}
	if divisor < 0 {
		divs = -divisor
	}

	if divs > divd {
		return 0
	}
	quotient := 0
	for divs <= divd {
		s, t := divs, 1
		for s+s <= divd {
			s += s
			t += t
		}
		divd -= s
		quotient += t
	}

	if negative {
		quotient = -quotient
	}
	return quotient
}
