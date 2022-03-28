package solution

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return res
	}

	for p1 := 0; p1 < n-2; p1++ {
		if p1 > 0 && nums[p1-1] == nums[p1] {
			continue
		}
		if nums[p1] > 0 {
			break
		}
		p2 := p1 + 1
		p3 := n - 1
		for p2 < p3 {
			if p2 > p1+1 && nums[p2-1] == nums[p2] {
				p2++
				continue
			}
			if p3 < n-1 && nums[p3] == nums[p3+1] {
				p3--
				continue
			}
			sum := nums[p1] + nums[p2] + nums[p3]
			if sum == 0 {
				res = append(res, []int{nums[p1], nums[p2], nums[p3]})
				p2++
				p3--
			} else if sum < 0 {
				p2++
			} else {
				p3--
			}
		}

	}

	return res
}
