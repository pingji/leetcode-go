package solution

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = i
	}

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			index2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if index1, ok := m[nums2[index2]]; ok {
				ans[index1] = nums2[i]
			}
		}
		stack = append(stack, i)
	}

	return ans
}
