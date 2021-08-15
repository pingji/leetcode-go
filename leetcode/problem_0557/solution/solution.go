package solution

func reverseWords(s string) string {
	arr := []byte(s)
	space := byte(' ')
	length := len(arr)
	i := 0

	for i < length {
		// 1. find a word
		left := i
		for i < length && arr[i] != space {
			i++
		}
		// 2. reverse a word
		right := i - 1
		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
		// 3. skip space
		for i < length && arr[i] == space {
			i++
		}
	}

	return string(arr)
}
