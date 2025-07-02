package quicksort2

import (
	"math/rand"
)

// partition 函数用于将数组分区，返回基准元素的位置
// 参数：nums - 待排序数组，l - 左边界，r - 右边界
// 返回值：基准元素在分区后的位置
// 使用双指针法进行分区，比单指针法更高效
func partition(nums []int, l int, r int) int {
	// 随机选择基准元素，避免最坏情况（已排序数组）
	x := rand.Intn(r-l+1) + l
	// 将基准元素交换到左边界位置
	nums[l], nums[x] = nums[x], nums[l]
	// 保存基准元素的值
	v := nums[l]

	// 双指针法分区：
	// [l+1, i) 区间存储小于等于基准元素的数
	// (j, r] 区间存储大于等于基准元素的数
	// [i, j] 区间是待处理的元素
	i := l + 1
	j := r
	for {
		// 从左向右找到第一个大于等于基准元素的元素
		for i <= j && nums[i] < v {
			i++
		}
		// 从右向左找到第一个小于等于基准元素的元素
		for i <= j && nums[j] > v {
			j--
		}
		// 如果左右指针相遇或交叉，分区完成
		if i > j {
			break
		}
		// 交换左右指针指向的元素，将它们放到正确的位置
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	// 将基准元素放到正确的位置（j位置）
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// quickSort 递归实现快速排序的核心函数
// 参数：nums - 待排序数组，l - 左边界，r - 右边界
func quickSort(nums []int, l int, r int) {
	// 递归终止条件：当左边界大于等于右边界时
	if l >= r {
		return
	}
	// 进行分区操作，获取基准元素位置
	p := partition(nums, l, r)
	// 递归排序左半部分（小于基准元素的部分）
	quickSort(nums, l, p-1)
	// 递归排序右半部分（大于基准元素的部分）
	quickSort(nums, p+1, r)
}

// QuickSort 快速排序的入口函数
// 参数：nums - 待排序的整数数组
// 时间复杂度：平均 O(nlogn)，最坏 O(n²)
// 空间复杂度：O(logn)，递归调用栈的深度
// 相比quicksort1，这个版本使用双指针法，在处理重复元素时更高效
func QuickSort(nums []int) {
	// 如果数组为空或只有一个元素，直接返回
	if len(nums) <= 1 {
		return
	}
	// 调用核心排序函数，传入数组和边界
	quickSort(nums, 0, len(nums)-1)
}
