package solution

// nextPermutation 下一个排列
// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
func nextPermutation(nums []int) {
	n := len(nums)

	// 第一步：从右向左找到第一个相邻的升序对 (nums[i] < nums[i+1])
	// 这个位置 i 就是需要交换的位置
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 第二步：如果找到了这样的位置 i
	if i >= 0 {
		// 从右向左找到第一个大于 nums[i] 的数 nums[j]
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 第三步：将 i+1 位置之后的所有数字反转
	// 这样可以保证得到的是下一个排列（字典序最小的更大排列）
	reverse(nums[i+1:])
}

// reverse 反转数组
// 使用双指针技术将数组从两端向中间反转
func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
