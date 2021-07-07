package solution

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = true
	}
	return -1
}
