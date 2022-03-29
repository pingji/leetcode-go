package solution

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	if n < 4 {
		return res
	}

	for p1 := 0; p1 < n-3; p1++ {
		if p1 > 0 && nums[p1-1] == nums[p1] {
			continue
		}
		for p2 := p1 + 1; p2 < n-2; p2++ {
			if p2 > p1+1 && nums[p2-1] == nums[p2] {
				continue
			}

			p3 := p2 + 1
			p4 := n - 1
			for p3 < p4 {
				if p3 > p2+1 && nums[p3-1] == nums[p3] {
					p3++
					continue
				}
				if p4 < n-1 && nums[p4] == nums[p4+1] {
					p4--
					continue
				}
				sum := nums[p1] + nums[p2] + nums[p3] + nums[p4]
				if sum == target {
					res = append(res, []int{nums[p1], nums[p2], nums[p3], nums[p4]})
					p3++
					p4--
				} else if sum < target {
					p3++
				} else {
					p4--
				}
			}
		}
	}

	return res
}
