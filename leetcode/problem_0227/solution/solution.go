package solution

func calculate(s string) int {
	ans := 0
	stack := []int{}
	// add a + before first digit
	preOpt := '+'

	num := 0
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		isSpece := ch == ' '
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if (!isDigit && !isSpece) || i == len(s)-1 {
			switch preOpt {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			preOpt = ch
			num = 0
		}
	}

	for _, v := range stack {
		ans += v
	}

	return ans
}
