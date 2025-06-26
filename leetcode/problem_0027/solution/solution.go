package solution

// removeElement 移除元素
// 给你一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，
// 并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func removeElement(nums []int, val int) int {
	// 处理空数组的情况
	if nums == nil {
		return 0
	}

	// 使用双指针技术
	// j 指向下一个要放置保留元素的位置
	j := 0

	// 遍历整个数组
	for i := 0; i < len(nums); i++ {
		// 如果当前元素不等于要删除的值
		if nums[i] != val {
			// 只有当 i 和 j 不同时才需要复制（避免不必要的自我赋值）
			if i != j {
				nums[j] = nums[i]
			}
			// 将 j 指针向前移动一位，为下一个保留元素准备位置
			j++
		}
		// 如果当前元素等于要删除的值，则跳过（不复制，不移动 j）
	}

	// 返回新数组的长度（j 就是保留元素的个数）
	return j
}
