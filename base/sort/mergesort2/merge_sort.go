package mergesort2

// merge 合并两个有序数组 [l, mid] 和 [mid+1, r]
// 使用临时数组来存储合并结果，避免覆盖原数组
func merge(nums []int, l int, mid int, r int) {
	// 创建临时数组，复制原数组的 [l, r] 区间
	temp := make([]int, r-l+1)
	copy(temp, nums[l:r+1])

	i := l       // 左半部分的指针
	j := mid + 1 // 右半部分的指针
	// 合并两个有序数组到原数组中
	for k := l; k <= r; k++ {
		if i > mid {
			// 左半部分已经遍历完，直接取右半部分的元素
			nums[k] = temp[j-l]
			j++
		} else if j > r {
			// 右半部分已经遍历完，直接取左半部分的元素
			nums[k] = temp[i-l]
			i++
		} else if temp[i-l] < temp[j-l] {
			// 左半部分的元素更小，取左半部分的元素
			nums[k] = temp[i-l]
			i++
		} else {
			// 右半部分的元素更小或相等，取右半部分的元素
			nums[k] = temp[j-l]
			j++
		}
	}
}

// min 返回两个整数中的较小值
func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

// MergeSort 迭代版本的归并排序
// 时间复杂度：O(n log n)，空间复杂度：O(n)
// 算法思想：自底向上，先合并相邻的两个元素，再合并相邻的四个元素，以此类推
func MergeSort(nums []int) {
	// step 表示当前要合并的子数组大小，从1开始，每次翻倍
	for step := 1; step <= len(nums); step += step {
		// 遍历所有需要合并的子数组对
		// i 是第一个子数组的起始位置，i+step 是第二个子数组的起始位置
		// i+step < len(nums) 确保第二个子数组存在，避免越界
		// i += 2*step 表示每次跳过两个子数组的距离，确保合并相邻的有序子数组
		for i := 0; i+step < len(nums); i += 2 * step {
			// 计算右边界，确保不越界
			var r = min(i+2*step-1, len(nums)-1)
			// 合并 [i, i+step-1] 和 [i+step, r] 两个有序子数组
			merge(nums, i, i+step-1, r)
		}
	}
}
