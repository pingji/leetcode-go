package solution

func findReplaceIndex(tail []int, num int) int {
	l, r := 0, len(tail)-1
	for l <= r {
		mid := (l + r) / 2
		if tail[mid] == num {
			return mid
		} else if tail[mid] < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tail := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
		} else {
			// bainary search
			index := findReplaceIndex(tail, nums[i])
			tail[index] = nums[i]
		}
	}

	return len(tail)
}
