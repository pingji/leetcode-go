package quicksort3

import (
	"math/rand"
	"time"
)

func partition(nums []int, l int, r int) (int, int) {
	x := rand.Intn(r-l+1) + l
	nums[l], nums[x] = nums[x], nums[l]
	v := nums[l]

	// [l+1, lt] < v
	// [lt+1, i) = v
	// [gt, r] > v

	lt := l
	gt := r + 1
	i := l + 1

	for i < gt {
		if nums[i] < v {
			nums[lt+1], nums[i] = nums[i], nums[lt+1]
			lt++
			i++
		} else if nums[i] == v {
			i++
		} else {
			nums[gt-1], nums[i] = nums[i], nums[gt-1]
			gt--
		}
	}

	nums[l], nums[lt] = nums[lt], nums[l]
	lt--
	return lt, gt
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	lt, gt := partition(nums, l, r)
	quickSort(nums, l, lt)
	quickSort(nums, gt, r)
}

func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
}
