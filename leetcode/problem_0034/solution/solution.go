package solution

func binarySearch(nums []int, target int, equal bool) int {
	left, right := 0, len(nums)-1
	ret := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target || (nums[mid] >= target && equal) {
			right = mid - 1
			ret = mid
		} else {
			left = mid + 1
		}
	}
	return ret
}

func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false) - 1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	} else {
		return []int{-1, -1}
	}
}
