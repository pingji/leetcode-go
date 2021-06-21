package solution

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 0
	cnt := make(map[byte]int)
	l := 0
	r := 0
	for ; r < len(s); r++ {
		cnt[s[r]]++
		maxLen = max(maxLen, cnt[s[r]])
		if r-l+1 > maxLen+k {
			cnt[s[l]]--
			l++
		}

	}

	return r - l
}
