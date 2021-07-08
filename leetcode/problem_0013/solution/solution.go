package solution

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	length := len(s)
	for i := 0; i < length; i++ {
		v := m[s[i]]
		if i < length-1 && v < m[s[i+1]] {
			sum -= v
		} else {
			sum += v
		}
	}

	return sum
}
