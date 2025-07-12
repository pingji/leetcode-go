package solution

import (
	"math/rand"
	// "time"
)

func partion(nums []int, l int, r int) int {
	x := rand.Intn(r-l+1) + l
	nums[l], nums[x] = nums[x], nums[l]
	v := nums[l]

	// [l+1, i) >= nums[l]
	// (j, r] <= nums[l]
	i := l + 1
	j := r
	for {
		for i <= j && nums[i] > v {
			i++
		}
		for i <= j && nums[j] < v {
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

func quickSearch(nums []int, l int, r int, k int) int {
	p := partion(nums, l, r)
	if p == k {
		return nums[p]
	} else if p < k {
		return quickSearch(nums, p+1, r, k)
	} else {
		return quickSearch(nums, l, p-1, k)
	}
}

func findKthLargest(nums []int, k int) int {
	// rand.Seed(time.Now().UnixNano())
	return quickSearch(nums, 0, len(nums)-1, k-1)
}
