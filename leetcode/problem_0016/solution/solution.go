package solution

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	best := math.MaxInt32
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return best
	}

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for p1 := 0; p1 < n-2; p1++ {
		if p1 > 0 && nums[p1-1] == nums[p1] {
			continue
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
			update(sum)
			if sum == target {
				return target
			} else if sum < target {
				p2++
			} else {
				p3--
			}
		}

	}

	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
