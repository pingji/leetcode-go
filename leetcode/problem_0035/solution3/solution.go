package solution

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
