package quicksort1

import (
	"math/rand"
	"time"
)

func partition(nums []int, l int, r int) int {
	x := rand.Intn(r-l+1) + l
	nums[l], nums[x] = nums[x], nums[l]

	// [l+1, j] < nums[l]
	// [j+1, i] >= nums[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < nums[l] {
			nums[j+1], nums[i] = nums[i], nums[j+1]
			j++
		}
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

func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
}
