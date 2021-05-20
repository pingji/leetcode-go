package solution

import (
	"math/rand"
	"time"
)

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

func quickSearch(arr []int, l int, r int, k int) []int {
	p := partition(arr, l, r)
	if p == k {
		return arr[:k+1]
	} else if p < k {
		return quickSearch(arr, p+1, r, k)
	} else {
		return quickSearch(arr, l, p-1, k)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	rand.Seed(time.Now().UnixNano())
	return quickSearch(arr, 0, len(arr)-1, k-1)
}
