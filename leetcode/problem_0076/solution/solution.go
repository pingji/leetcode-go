package solution

func minWindow(s string, t string) string {
	if s == t {
		return s
	}
	if len(s) < len(t) {
		return ""
	}

	sCnt := make(map[byte]int)
	tCnt := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tCnt[t[i]]++
	}

	check := func() bool {
		for k, v := range tCnt {
			if sCnt[k] < v {
				return false
			}
		}
		return true
	}

	resL := -1
	resR := -1
	minLen := len(s) + 1

	for l, r := 0, 0; r < len(s); r++ {
		if tCnt[s[r]] > 0 {
			sCnt[s[r]]++
		}
		for ; check() && l <= r; l++ {
			sCnt[s[l]]--
			sLen := r - l + 1
			if sLen < minLen {
				minLen = r - l + 1
				resL = l
				resR = r
			}
		}

	}

	if resL == -1 {
		return ""
	}

	return s[resL : resR+1]
}
