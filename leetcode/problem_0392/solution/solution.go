package solution

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	ps := 0
	pt := 0
	for ps < len(s) && pt < len(t) {
		if s[ps] == t[pt] {
			ps++
		}
		pt++
		if ps == len(s) {
			return true
		}
	}
	return false
}
