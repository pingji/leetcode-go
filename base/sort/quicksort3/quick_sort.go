package quicksort3

import (
	"math/rand"
)

// partition 三路分区函数，将数组分为三部分：小于、等于、大于基准元素
// 参数：nums - 待排序数组，l - 左边界，r - 右边界
// 返回值：(lt, gt) - 等于基准元素的区间边界
// 三路快排专门处理包含大量重复元素的数组，避免重复元素导致的性能退化
func partition(nums []int, l int, r int) (int, int) {
	// 随机选择基准元素，避免最坏情况（已排序数组）
	x := rand.Intn(r-l+1) + l
	// 将基准元素交换到左边界位置
	nums[l], nums[x] = nums[x], nums[l]
	// 保存基准元素的值
	v := nums[l]

	// 三路分区：
	// [l+1, lt] 区间存储小于基准元素的数
	// [lt+1, i) 区间存储等于基准元素的数
	// [i, gt) 区间是待处理的元素
	// [gt, r] 区间存储大于基准元素的数
	lt := l     // 小于区间的右边界
	gt := r + 1 // 大于区间的左边界
	i := l + 1  // 当前处理的位置

	for i < gt {
		if nums[i] < v {
			// 当前元素小于基准元素，交换到小于区间
			nums[lt+1], nums[i] = nums[i], nums[lt+1]
			lt++
			i++
		} else if nums[i] == v {
			// 当前元素等于基准元素，直接跳过（已在等于区间）
			i++
		} else {
			// 当前元素大于基准元素，交换到大于区间
			nums[gt-1], nums[i] = nums[i], nums[gt-1]
			gt--
		}
	}

	// 将基准元素放到正确的位置（等于区间的开始位置）
	nums[l], nums[lt] = nums[lt], nums[l]
	lt-- // 调整lt位置，现在lt指向小于区间的最后一个元素
	return lt, gt
}

// quickSort 递归实现三路快速排序的核心函数
// 参数：nums - 待排序数组，l - 左边界，r - 右边界
func quickSort(nums []int, l int, r int) {
	// 递归终止条件：当左边界大于等于右边界时
	if l >= r {
		return
	}
	// 进行三路分区操作，获取小于和大于区间的边界
	lt, gt := partition(nums, l, r)
	// 递归排序小于基准元素的部分
	quickSort(nums, l, lt)
	// 递归排序大于基准元素的部分
	// 注意：等于基准元素的部分已经在正确位置，无需排序
	quickSort(nums, gt, r)
}

// QuickSort 三路快速排序的入口函数
// 参数：nums - 待排序的整数数组
// 时间复杂度：平均 O(nlogn)，最坏 O(n²)
// 空间复杂度：O(logn)，递归调用栈的深度
// 优势：对于包含大量重复元素的数组，性能显著优于普通快排
// 避免了重复元素导致的性能退化问题
func QuickSort(nums []int) {
	// 如果数组为空或只有一个元素，直接返回
	if len(nums) <= 1 {
		return
	}

	// 调用核心排序函数，传入数组和边界
	quickSort(nums, 0, len(nums)-1)
}
