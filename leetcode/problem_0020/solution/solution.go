package solution

// isValid 判断字符串中的括号是否有效匹配
// 使用栈来匹配括号对，遇到左括号入栈，遇到右括号与栈顶匹配
func isValid(s string) bool {
	n := len(s)
	// 如果字符串长度为奇数，必定无法匹配
	if n%2 == 1 {
		return false
	}

	// 定义右括号到左括号的映射关系
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 使用栈存储左括号
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		// 检查当前字符是否为右括号
		if _, ok := pairs[s[i]]; ok {
			// 如果是右括号，检查栈是否为空或栈顶是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 匹配成功，弹出栈顶的左括号
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，直接入栈
			stack = append(stack, s[i])
		}
	}

	// 最后检查栈是否为空，为空说明所有括号都匹配成功
	return len(stack) == 0
}
