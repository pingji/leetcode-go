package solution

// merge sort

func merge(nums []int, l int, mid int, r int) {
	temp := make([]int, r-l+1)
	copy(temp, nums[l:r+1])

	i := 0
	j := mid + 1 - l
	for k := l; k <= r; k++ {
		if i > mid-l {
			nums[k] = temp[j]
			j++
		} else if j > r-l {
			nums[k] = temp[i]
			i++
		} else if temp[i] < temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
		}
	}
}

func mergeSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r)
	}
}

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
