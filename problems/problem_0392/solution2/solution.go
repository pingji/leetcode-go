package solution

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	// init f[i][j]
	f := make([][26]int, m+1)
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}
	// search
	k := 0
	for i := 0; i < n; i++ {
		if f[k][int(s[i]-'a')] == m {
			return false
		}
		k = f[k][int(s[i]-'a')] + 1
	}

	return true
}
