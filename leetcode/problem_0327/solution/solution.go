package solution

// merge 函数使用归并排序的思想来计算满足条件的区间和个数
// 参数说明：
// - sum: 前缀和数组
// - lower, upper: 区间和的范围限制
// - left, right: 当前处理的数组范围
// 返回值：满足条件的区间和个数
func merge(sum []int, lower int, upper int, left int, right int) int {
	// 基本情况：如果只有一个元素，无法形成区间，返回0
	if left == right {
		return 0
	}

	// 计算中点，将数组分为左右两部分
	mid := (right + left) / 2

	// 递归处理左右两个子数组，并累加它们的计数结果
	n1 := merge(sum, lower, upper, left, mid)
	n2 := merge(sum, lower, upper, mid+1, right)
	count := n1 + n2

	// 使用双指针技巧在右半部分寻找满足条件的范围
	// l 指向第一个满足 sum[l] - sum[i] >= lower 的位置
	// r 指向第一个满足 sum[r] - sum[i] > upper 的位置
	l := mid + 1
	r := mid + 1

	// 对于左半部分的每个元素 sum[i]，在右半部分寻找满足条件的区间
	for i := left; i <= mid; i++ {
		// 找到第一个满足 sum[l] - sum[i] >= lower 的位置
		for l <= right && sum[l]-sum[i] < lower {
			l++
		}
		// 找到第一个满足 sum[r] - sum[i] > upper 的位置
		for r <= right && sum[r]-sum[i] <= upper {
			r++
		}
		// 满足条件的区间个数为 r - l
		count += (r - l)
	}

	// 归并排序：合并两个有序子数组
	// 创建临时数组来存储当前范围的元素
	temp := make([]int, right-left+1)
	copy(temp, sum[left:right+1])

	// 计算临时数组中的索引
	tmpMid := mid - left
	tmpRight := right - left
	p1 := 0              // 左半部分的指针
	p2 := mid + 1 - left // 右半部分的指针

	// 合并两个有序子数组
	for k := left; k <= right; k++ {
		if p1 > tmpMid {
			// 左半部分已经处理完，直接取右半部分的元素
			sum[k] = temp[p2]
			p2++
		} else if p2 > tmpRight {
			// 右半部分已经处理完，直接取左半部分的元素
			sum[k] = temp[p1]
			p1++
		} else if temp[p1] < temp[p2] {
			// 左半部分的元素较小，取左半部分的元素
			sum[k] = temp[p1]
			p1++
		} else {
			// 右半部分的元素较小或相等，取右半部分的元素
			sum[k] = temp[p2]
			p2++
		}
	}
	return count
}

// countRangeSum 计算满足条件的区间和个数
// 参数说明：
// - nums: 输入数组
// - lower, upper: 区间和的范围限制
// 返回值：满足条件的区间和个数
func countRangeSum(nums []int, lower int, upper int) int {
	// 计算前缀和数组
	// sum[0] = 0 (空数组的前缀和)
	// sum[i+1] 表示前 i 个元素的和，即 nums[0] 到 nums[i] 的和
	sum := make([]int, len(nums)+1)
	sum[0] = 0 // 前缀和的起始值
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	// 调用归并排序函数来计算满足条件的区间个数
	return merge(sum, lower, upper, 0, len(sum)-1)
}
