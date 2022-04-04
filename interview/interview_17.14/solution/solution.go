package solution

import (
	"math/rand"
	"time"
)

func partition(arr []int, l, r int) int {
	x := rand.Intn(r-l+1) + l
	arr[l], arr[x] = arr[x], arr[l]
	v := arr[l]

	// [l+1, i) <= arr[l]
	// (j, r] >= arr[l]
	i := l + 1
	j := r
	for {
		for i <= j && arr[i] < v {
			i++
		}
		for i <= j && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func smallestK(arr []int, k int) []int {
	rand.Seed(time.Now().UnixNano())
	findSmallestK(arr, 0, len(arr)-1, k-1)
	return arr[:k]
}

func findSmallestK(arr []int, l int, r int, k int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	if p == k {
		return
	} else if p < k {
		findSmallestK(arr, p+1, r, k)
	} else {
		findSmallestK(arr, l, p-1, k)
	}
}
