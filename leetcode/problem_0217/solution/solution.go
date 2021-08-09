package solution

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		} else {
			set[v] = true
		}
	}
	return false
}
