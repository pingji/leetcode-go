package solution

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen < 2 {
		return s
	}

	maxStart := 0
	maxEnd := 0
	maxLen := 0

	for i := 0; i < strLen; i++ {
		left := i - 1
		right := i + 1
		curLen := 1
		for left >= 0 && s[left] == s[i] {
			curLen += 1
			left--
		}
		for right < strLen && s[right] == s[i] {
			curLen += 1
			right++
		}
		for left >= 0 && right < strLen && s[left] == s[right] {
			curLen += 2
			left--
			right++
		}
		if curLen > maxLen {
			maxLen = curLen
			maxStart = left + 1
			maxEnd = right - 1
		}
	}

	return s[maxStart : maxEnd+1]
}
