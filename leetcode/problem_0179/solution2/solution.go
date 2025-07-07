package solution

import (
	"sort"
	"strconv"
	"strings"
)

// largestNumber 将非负整数数组重新排列，组成最大的数
// 参数：nums - 非负整数数组
// 返回值：重新排列后组成的最大数的字符串表示
// 算法思路：
// 1. 将所有数字转换为字符串
// 2. 自定义排序规则：比较两个字符串拼接后的字典序
// 3. 按自定义规则排序后拼接所有字符串
// 4. 处理前导零的特殊情况
// 时间复杂度：O(nlogn * k)，其中k是数字的平均位数
// 空间复杂度：O(n)，存储字符串数组
func largestNumber(nums []int) string {
	// 将所有数字转换为字符串数组
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}

	// 自定义排序规则：比较两个字符串拼接后的字典序
	// 对于字符串 a 和 b，如果 a+b > b+a，则 a 应该排在 b 前面
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})

	// 将所有排序后的字符串拼接
	s := strings.Join(ss, "")

	// 处理前导零的特殊情况
	// 如果排序后的第一个字符是'0'，说明所有数字都是0，返回"0"
	if s[0] == '0' {
		return "0"
	}
	return s
}
