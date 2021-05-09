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

	var recursion func(index int, path string)
	recursion = func(index int, path string) {
		if len(digits) == len(path) {
			res = append(res, path)
			return
		}
		d := string(digits[index])
		str := photoMap[d]
		for i := range str {
			recursion(index+1, path+string(str[i]))
		}
	}
	recursion(0, "")
	return res
}
