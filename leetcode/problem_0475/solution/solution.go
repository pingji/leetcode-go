package solution

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	maxDist := 0
	for _, house := range houses {
		left := binarySearch(heaters, house)
		right := left + 1
		leftDist, rightDist, minDist := math.MaxInt32, math.MaxInt32, math.MaxInt32
		if left >= 0 {
			leftDist = house - heaters[left]
		}
		if right <= len(heaters)-1 {
			rightDist = heaters[right] - house
		}
		if leftDist <= rightDist {
			minDist = leftDist
		} else {
			minDist = rightDist
		}
		if minDist > maxDist {
			maxDist = minDist
		}
	}
	return maxDist
}

func binarySearch(heaters []int, target int) int {
	left, right := 0, len(heaters)-1
	if heaters[left] > target {
		return -1
	}
	for left < right {
		mid := (right + left + 1) / 2
		if heaters[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
