package solution

import "sort"

// threeSum 三数之和
// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
func threeSum(nums []int) [][]int {
	res := [][]int{}
	// 先对数组进行排序，便于去重和双指针操作
	sort.Ints(nums)
	n := len(nums)

	// 如果数组长度小于3，直接返回空结果
	if n < 3 {
		return res
	}

	// 固定第一个数，用双指针寻找另外两个数
	for p1 := 0; p1 < n-2; p1++ {
		// 跳过重复的第一个数，避免重复结果
		if p1 > 0 && nums[p1-1] == nums[p1] {
			continue
		}
		// 如果第一个数大于0，由于数组已排序，后面不可能有和为0的三元组
		if nums[p1] > 0 {
			break
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

			if sum == 0 {
				// 找到满足条件的三元组
				res = append(res, []int{nums[p1], nums[p2], nums[p3]})
				p2++ // 移动左指针
				p3-- // 移动右指针
			} else if sum < 0 {
				// 和小于0，需要增大，移动左指针
				p2++
			} else {
				// 和大于0，需要减小，移动右指针
				p3--
			}
		}
	}

	return res
}
