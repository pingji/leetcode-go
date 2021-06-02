package solution

func mergeSort(nums []int, left int, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	n1 := mergeSort(nums, left, mid)
	n2 := mergeSort(nums, mid+1, right)
	count := n1 + n2

	tmp := []int{}
	i := left
	j := mid + 1

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			count += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		count += right - (mid + 1) + 1
		i++
	}
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}
	for i := left; i <= right; i++ {
		nums[i] = tmp[i-left]
	}
	return count
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}
