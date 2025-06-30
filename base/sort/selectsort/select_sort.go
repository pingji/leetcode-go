package selectsort

// SelectSort 选择排序算法
// 时间复杂度：O(n²)，空间复杂度：O(1)
// 算法思想：每次从未排序区间选择最小的元素，放到已排序区间的末尾
func SelectSort(nums []int) {
	// 遍历数组的每个位置，将最小元素放到当前位置
	for i := 0; i < len(nums); i++ {
		minIndex := i // 假设当前位置的元素是最小的
		// 在未排序区间中寻找最小元素
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j // 更新最小元素的索引
			}
		}
		// 如果找到的最小元素不在当前位置，则交换
		if i != minIndex {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}
