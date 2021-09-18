package solution

func findLongestSubString(s string) string {
	arr := []byte(s)
	var res []byte

	fast := 0
	slow := 0
	tempMaxLen := 0
	maxLen := 0

	for fast < len(arr) {
		tempMaxLen++

		if arr[fast] != arr[slow] || fast == len(arr)-1 {
			if tempMaxLen > maxLen {
				maxLen = tempMaxLen
				res = arr[slow:fast]
			}
			slow = fast
			tempMaxLen = 0
		}
		fast++
	}

	return string(res)
}
