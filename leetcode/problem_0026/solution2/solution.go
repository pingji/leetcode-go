package solution

// removeDuplicates 删除有序数组中的重复项
// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，
// 返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates(nums []int) int {
	// 如果数组长度小于等于1，直接返回长度
	if len(nums) <= 1 {
		return len(nums)
	}

	// 使用双指针技术
	// slow 指针指向当前不重复序列的最后一个位置
	slow := 0

	// fast 指针遍历整个数组，寻找与 slow 指向元素不同的元素
	for fast := 1; fast < len(nums); fast++ {
		// 当 fast 指向的元素与 slow 指向的元素不同时
		if nums[slow] != nums[fast] {
			// 将 slow 指针向前移动一位
			slow += 1
			// 将 fast 指向的元素复制到 slow 的新位置
			nums[slow] = nums[fast]
		}
		// 如果 fast 指向的元素与 slow 指向的元素相同，则跳过（删除重复项）
	}

	// 返回新数组的长度（slow + 1，因为 slow 是索引）
	return slow + 1
}
