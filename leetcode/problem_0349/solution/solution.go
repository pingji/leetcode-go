package solution

func intersection(nums1 []int, nums2 []int) []int {
	ans := []int{}
	set1 := make(map[int]struct{})
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := make(map[int]struct{})
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	for v := range set1 {
		if _, has := set2[v]; has {
			ans = append(ans, v)
		}
	}

	return ans
}
