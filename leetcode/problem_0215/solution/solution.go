package solution

import (
	"math/rand"
)

// partion 函数用于将数组分区，返回基准元素的位置
// 参数：nums - 待排序数组，l - 左边界，r - 右边界
// 返回值：基准元素在分区后的位置
// 使用双指针法进行分区，专门用于寻找第k大元素
func partion(nums []int, l int, r int) int {
	// 随机选择基准元素，避免最坏情况（已排序数组）
	x := rand.Intn(r-l+1) + l
	// 将基准元素交换到左边界位置
	nums[l], nums[x] = nums[x], nums[l]
	// 保存基准元素的值
	v := nums[l]

	// 双指针法分区：
	// [l+1, i) 区间存储大于等于基准元素的数（第k大元素在左侧）
	// (j, r] 区间存储小于等于基准元素的数（第k大元素在右侧）
	// [i, j] 区间是待处理的元素
	i := l + 1
	j := r
	for {
		// 从左向右找到第一个小于等于基准元素的元素
		for i <= j && nums[i] > v {
			i++
		}
		// 从右向左找到第一个大于等于基准元素的元素
		for i <= j && nums[j] < v {
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

// quickSearch 递归实现快速选择算法的核心函数
// 参数：nums - 待搜索数组，l - 左边界，r - 右边界，k - 目标位置
// 返回值：第k大元素的值
// 时间复杂度：平均 O(n)，最坏 O(n²)
// 空间复杂度：O(logn)，递归调用栈的深度
func quickSearch(nums []int, l int, r int, k int) int {
	// 进行分区操作，获取基准元素位置
	p := partion(nums, l, r)
	// 如果基准元素位置等于目标位置k，找到第k大元素
	if p == k {
		return nums[p]
	} else if p < k {
		// 如果基准元素位置小于目标位置k，在右半部分继续搜索
		return quickSearch(nums, p+1, r, k)
	} else {
		// 如果基准元素位置大于目标位置k，在左半部分继续搜索
		return quickSearch(nums, l, p-1, k)
	}
}

// findKthLargest 寻找数组中第k大的元素
// 参数：nums - 整数数组，k - 第k大（从1开始计数）
// 返回值：第k大元素的值
// 算法：使用快速选择算法，基于快速排序的分区思想
// 优势：平均时间复杂度O(n)，比排序后取第k个元素更高效
func findKthLargest(nums []int, k int) int {
	// 将k转换为从0开始的索引（因为数组索引从0开始）
	return quickSearch(nums, 0, len(nums)-1, k-1)
}
