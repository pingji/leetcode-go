package solution

import "bytes"

func frequencySort(s string) string {
	cnt := make(map[byte]int)
	maxFreq := 0
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for c, n := range cnt {
		buckets[n] = append(buckets[n], c)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, c := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{c}, i)...)
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
