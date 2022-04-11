package solution

import "math"

var (
	states = map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	sign  = 1
	ans   = 0
	state = "start"
)

func getCol(c byte) int {
	if c == ' ' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}

func handleChar(c byte) {
	col := getCol(c)
	state = states[state][col]
	if state == "signed" {
		if c == '-' {
			sign = -1
		}
	}
	if state == "in_number" {
		ans = ans*10 + int(c-'0')
		if sign == 1 {
			ans = min(ans, math.MaxInt32)
		} else {
			ans = min(ans, -math.MinInt32)
		}
	}

}

func myAtoi(s string) int {
	sign = 1
	ans = 0
	state = "start"
	for i := 0; i < len(s); i++ {
		if state == "end" {
			break
		}
		handleChar(s[i])
	}
	if sign == 1 {
		return ans
	} else {
		return -ans
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
