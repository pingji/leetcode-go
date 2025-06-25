package solution

import (
	"math"
	"sort"
)

// threeSumClosest 最接近的三数之和
// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
// 找出 nums 中的三个整数，使得它们的和与 target 最接近。
// 返回这三个数的和。假定每组输入只存在唯一答案。
func threeSumClosest(nums []int, target int) int {
	// 初始化最佳结果为最大整数值
	best := math.MaxInt32
	// 先对数组进行排序，便于双指针操作
	sort.Ints(nums)
	n := len(nums)

	// 如果数组长度小于3，返回初始值
	if n < 3 {
		return best
	}

	// 更新最佳结果的辅助函数
	// 如果当前和与目标值的距离小于最佳结果与目标值的距离，则更新最佳结果
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 固定第一个数，用双指针寻找另外两个数
	for p1 := 0; p1 < n-2; p1++ {
		// 跳过重复的第一个数，避免重复计算
		if p1 > 0 && nums[p1-1] == nums[p1] {
			continue
		}

		// 使用双指针寻找另外两个数
		p2 := p1 + 1 // 左指针
		p3 := n - 1  // 右指针

		for p2 < p3 {
			// 跳过重复的第二个数
			if p2 > p1+1 && nums[p2-1] == nums[p2] {
				p2++
				continue
			}
			// 跳过重复的第三个数
			if p3 < n-1 && nums[p3] == nums[p3+1] {
				p3--
				continue
			}

			// 计算三数之和
			sum := nums[p1] + nums[p2] + nums[p3]
			// 更新最佳结果
			update(sum)

			if sum == target {
				// 如果找到精确匹配，直接返回目标值
				return target
			} else if sum < target {
				// 和小于目标值，需要增大，移动左指针
				p2++
			} else {
				// 和大于目标值，需要减小，移动右指针
				p3--
			}
		}
	}

	return best
}

// abs 计算整数的绝对值
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
