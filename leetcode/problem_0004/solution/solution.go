package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		return float64(findKthElement(nums1, nums2, total/2+1))
	} else {
		return float64(findKthElement(nums1, nums2, total/2)+findKthElement(nums1, nums2, total/2+1)) / 2.0
	}
}

func findKthElement(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	index1 := min(k/2-1, len(nums1)-1)
	index2 := min(k/2-1, len(nums2)-1)
	if nums1[index1] <= nums2[index2] {
		k -= index1 + 1
		return findKthElement(nums1[index1+1:], nums2, k)
	} else {
		k -= index2 + 1
		return findKthElement(nums1, nums2[index2+1:], k)
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
