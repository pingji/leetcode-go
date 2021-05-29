package solution

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i int, j int) bool {
		a := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
		b := strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
		return a > b
	})

	var b strings.Builder
	for _, num := range nums {
		b.WriteString(strconv.Itoa(num))
	}

	s := b.String()
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		s = "0"
	}
	return s
}
