package solution

import (
	"math/rand"
	"time"
)

func partion(arr [][2]int, l int, r int) int {
	x := rand.Intn(r-l+1) + l
	arr[l], arr[x] = arr[x], arr[l]
	v := arr[l][1]

	// [l+1, i) >= arr[l][1]
	// (j, r] <= arr[l][1]
	i := l + 1
	j := r

	for {
		for i <= j && arr[i][1] > v {
			i++
		}
		for i <= j && arr[j][1] < v {
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

func quickSearch(arr [][2]int, l int, r int, k int) [][2]int {
	p := partion(arr, l, r)
	if p == k {
		return arr[:k+1]
	} else if p < k {
		return quickSearch(arr, p+1, r, k)
	} else {
		return quickSearch(arr, l, p-1, k)
	}
}

func topKFrequent(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}

	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	countArr := make([][2]int, len(countMap))
	index := 0
	for k, v := range countMap {
		countArr[index] = [2]int{k, v}
		index++
	}

	rand.Seed(time.Now().UnixNano())
	resArr := quickSearch(countArr, 0, len(countArr)-1, k-1)

	res := make([]int, len(resArr))
	for i := 0; i < len(resArr); i++ {
		res[i] = resArr[i][0]
	}
	return res
}
