package solution

func intersect(nums1 []int, nums2 []int) []int {
	ans := []int{}
	set1 := make(map[int]int)
	for _, v := range nums1 {
		if count, has := set1[v]; has {
			set1[v] = count + 1
		} else {
			set1[v] = 1
		}
	}

	for _, v := range nums2 {
		if count, has := set1[v]; has {
			if count > 0 {
				ans = append(ans, v)
				set1[v] = count - 1
			}
		}
	}

	return ans
}
