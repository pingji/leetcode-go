package solution

// removeElement 移除元素
// 给你一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，
// 并返回移除后数组的新长度。
func removeElement(nums []int, val int) int {
	// 处理空数组的情况
	if len(nums) == 0 {
		return 0
	}

	// 使用双指针技术
	// i 指向当前要检查的位置
	// j 指向数组末尾，用于存放要删除的元素
	i, j := 0, len(nums)-1

	for i <= j {
		if nums[i] == val {
			// 如果当前元素等于目标值，将其与末尾元素交换
			nums[i] = nums[j]
			j-- // 缩小数组范围
		} else {
			// 如果当前元素不等于目标值，保留该元素，检查下一个
			i++
		}
	}

	// 返回新数组的长度
	// j + 1 是因为 j 指向最后一个保留元素的索引
	return j + 1
}
