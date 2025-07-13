package solution

import (
	"math/rand"
	"time"
)

// quick sort
func partition(nums []int, l int, r int) int {
	x := rand.Intn(r-l+1) + l
	nums[l], nums[x] = nums[x], nums[l]
	v := nums[l]

	// [l+1, i) <= nums[l]
	// (j, r] >= nums[l]
	i := l + 1
	j := r
	for {
		for i <= j && nums[i] < v {
			i++
		}
		for i <= j && nums[j] > v {
			j--
		}
		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}
