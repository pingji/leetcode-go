package solution

import (
	"strings"
)

func reverseWords(s string) string {
	var revArr []string
	// 1. split
	arr := strings.Fields(s)
	// 2. reverse
	for i := len(arr) - 1; i >= 0; i-- {
		revArr = append(revArr, arr[i])
	}
	// 3. jion
	return strings.Join(revArr, " ")
}
