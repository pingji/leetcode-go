package solution

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		mid := (left + right) / 2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right -= 1
		}
	}
	return numbers[left]
}
