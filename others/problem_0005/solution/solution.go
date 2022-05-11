package solution

func secondFreqChar(s string) string {
	cnt := make(map[byte]int)
	maxFreq := 0
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}
	secondFreq := 0
	for _, n := range cnt {
		if n == maxFreq {
			continue
		}
		secondFreq = max(secondFreq, n)
	}
	ans := make([]byte, 0)
	for c, n := range cnt {
		if n == secondFreq {
			ans = append(ans, c)
		}
	}
	return string(ans)
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}
