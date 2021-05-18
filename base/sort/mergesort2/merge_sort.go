package mergesort2

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

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MergeSort(nums []int) {
	for step := 1; step <= len(nums); step += step {
		for i := 0; i+step < len(nums); i += 2 * step {
			var r = min(i+2*step-1, len(nums)-1)
			merge(nums, i, i+step-1, r)
		}
	}
}
