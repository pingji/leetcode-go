package solution

// lengthOfLongestSubstring 找到字符串中不包含重复字符的最长子串
// 使用滑动窗口算法，时间复杂度 O(n)，空间复杂度 O(min(m,n))，其中 m 是字符集大小
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := map[byte]int{} // 哈希表，记录每个字符在当前窗口中的出现次数
	count := 0          // 记录最长无重复字符子串的长度

	r := 0 // 右指针，指向当前窗口的右边界
	// 左指针 l 遍历字符串的每个位置
	for l := 0; l < len(s); l++ {
		// 如果不是第一个位置，需要将左指针前一个位置的字符从窗口中移除
		if l != 0 {
			delete(m, s[l-1]) // 从哈希表中完全删除该字符，而不是减少计数
		}
		// 扩展右指针，直到遇到重复字符或到达字符串末尾
		for r < len(s) && m[s[r]] == 0 {
			m[s[r]] = 1 // 将当前字符加入窗口
			r++         // 右指针右移
		}
		// 更新最长无重复字符子串的长度
		count = max(count, r-l)
		// 如果右指针已经到达字符串末尾，说明后续不会有更长的无重复子串
		if r == len(s) {
			break
		}
	}

	return count
}

// max 返回两个整数中的较大值
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
