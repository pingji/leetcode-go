package mergesort1

func merge(nums []int, l int, mid int, r int) {
	temp := make([]int, r-l+1)
	copy(temp, nums[l:r+1])

	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = temp[j-l]
			j++
		} else if j > r {
			nums[k] = temp[i-l]
			i++
		} else if temp[i-l] < temp[j-l] {
			nums[k] = temp[i-l]
			i++
		} else {
			nums[k] = temp[j-l]
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
	merge(nums, l, mid, r)
}

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}
