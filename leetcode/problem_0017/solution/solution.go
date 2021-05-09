package solution

var photoMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	path := make([]byte, 0, len(digits))

	var backtrack func(index int)
	backtrack = func(index int) {
		if len(digits) == len(path) {
			res = append(res, string(path))
			return
		}
		d := string(digits[index])
		str := photoMap[d]
		for i := range str {
			path = append(path, str[i])

			backtrack(index + 1)

			path = path[:len(path)-1]

		}
	}
	backtrack(0)
	return res
}
