package solution

func reverseStr(s string, k int) string {
	arr := []byte(s)
	length := len(s)

	i := 0
	for i < length {
		left := i
		for i < length && i < left+k {
			i++
		}
		right := i - 1

		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}

		left = i
		for i < length && i < left+k {
			i++
		}
	}

	return string(arr)
}
