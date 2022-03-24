package solution2

func findFirstPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			// next search range [left, mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// next search range [left, mid]
			right = mid
		} else {
			// nums[mid] < target
			// next search range [left+1, mid]
			left = mid + 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findLastPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] > target {
			// next search range [left, mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// next search range [mid, right]
			left = mid
		} else {
			// nums[mid] < target
			// next search range [left+1, mid]
			left = mid + 1
		}
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	left := findFirstPosition(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findLastPosition(nums, target)
	return []int{left, right}
}
