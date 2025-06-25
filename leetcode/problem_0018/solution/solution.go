package solution

import "sort"

// fourSum 四数之和
// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，
// 判断 nums 中是否存在四个元素 a，b，c 和 d ，
// 使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	// 先对数组进行排序，便于去重和双指针操作
	sort.Ints(nums)
	n := len(nums)

	// 如果数组长度小于4，直接返回空结果
	if n < 4 {
		return res
	}

	// 固定前两个数，用双指针寻找后两个数
	for p1 := 0; p1 < n-3; p1++ {
		// 跳过重复的第一个数，避免重复结果
		if p1 > 0 && nums[p1-1] == nums[p1] {
			continue
		}

		// 固定第二个数
		for p2 := p1 + 1; p2 < n-2; p2++ {
			// 跳过重复的第二个数，避免重复结果
			if p2 > p1+1 && nums[p2-1] == nums[p2] {
				continue
			}

			// 使用双指针寻找后两个数
			p3 := p2 + 1 // 左指针
			p4 := n - 1  // 右指针

			for p3 < p4 {
				// 跳过重复的第三个数
				if p3 > p2+1 && nums[p3-1] == nums[p3] {
					p3++
					continue
				}
				// 跳过重复的第四个数
				if p4 < n-1 && nums[p4] == nums[p4+1] {
					p4--
					continue
				}

				// 计算四数之和
				sum := nums[p1] + nums[p2] + nums[p3] + nums[p4]

				if sum == target {
					// 找到满足条件的四元组
					res = append(res, []int{nums[p1], nums[p2], nums[p3], nums[p4]})
					p3++ // 移动左指针
					p4-- // 移动右指针
				} else if sum < target {
					// 和小于目标值，需要增大，移动左指针
					p3++
				} else {
					// 和大于目标值，需要减小，移动右指针
					p4--
				}
			}
		}
	}

	return res
}
