package solution

// twoSum 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，
// 并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	// 使用哈希表存储已经遍历过的数字及其下标
	// key: 数字值, value: 数字在数组中的下标
	m := map[int]int{}

	// 遍历数组中的每个数字
	for i, num := range nums {
		// 计算需要寻找的补数
		complement := target - num

		// 检查补数是否已经在哈希表中
		if j, ok := m[complement]; ok {
			// 找到满足条件的两个数，返回它们的下标
			// j 是之前存储的补数的下标，i 是当前数字的下标
			return []int{j, i}
		}

		// 将当前数字及其下标存储到哈希表中，供后续查找使用
		m[num] = i
	}

	// 如果没有找到满足条件的两个数，返回 nil
	return nil
}
