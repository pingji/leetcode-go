package solution

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]int{}
	count := 0

	r := 0
	for l := 0; l < len(s); l++ {
		if l != 0 {
			m[s[l-1]]--
		}
		for r < len(s) && m[s[r]] == 0 {
			m[s[r]]++
			r++
		}
		count = max(count, r-l)
	}

	return count
}
