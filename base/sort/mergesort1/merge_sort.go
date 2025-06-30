package mergesort1

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

// mergeSort 递归实现归并排序
// 将数组 [l, r] 区间进行归并排序
func mergeSort(nums []int, l int, r int) {
	if l >= r {
		return // 基本情况：只有一个元素或空区间，已经有序
	}
	mid := (l + r) / 2        // 计算中点
	mergeSort(nums, l, mid)   // 递归排序左半部分
	mergeSort(nums, mid+1, r) // 递归排序右半部分
	// 优化：只有当左半部分的最大值大于右半部分的最小值时才需要合并
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r) // 合并两个有序数组
	}
}

// MergeSort 归并排序主函数
// 时间复杂度：O(n log n)，空间复杂度：O(n)
// 算法思想：分治法，将数组分成两半，分别排序后合并
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}
