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
// 1. 自定义排序规则：比较两个数字拼接后的字典序
// 2. 按自定义规则排序后拼接所有数字
// 3. 处理前导零的特殊情况
// 时间复杂度：O(nlogn * k)，其中k是数字的平均位数
// 空间复杂度：O(n)，存储字符串结果
func largestNumber(nums []int) string {
	// 自定义排序规则
	// 比较两个数字拼接后的字典序，选择字典序较大的排列方式
	sort.Slice(nums, func(i int, j int) bool {
		// 将两个数字转换为字符串并拼接
		a := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
		b := strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
		// 返回字典序较大的排列方式
		return a > b
	})

	// 使用 strings.Builder 高效拼接字符串
	var b strings.Builder
	for _, num := range nums {
		b.WriteString(strconv.Itoa(num))
	}

	// 获取拼接后的字符串
	s := b.String()
	// 去除前导零
	s = strings.TrimLeft(s, "0")
	// 如果去除前导零后字符串为空（说明原数组全为0），返回"0"
	if len(s) == 0 {
		s = "0"
	}
	return s
}
