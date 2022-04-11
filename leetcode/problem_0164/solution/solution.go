package solution

type Pair struct {
	min  int
	max  int
	size int
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	minVal, maxVal := findMinMax(nums)

	d := max(1, (maxVal-minVal)/(n-1))
	bucketSize := (maxVal-minVal)/d + 1
	buckets := make([]Pair, bucketSize)
	for _, v := range nums {
		id := (v - minVal) / d
		if buckets[id].size == 0 || v < buckets[id].min {
			buckets[id].min = v
		}
		if buckets[id].size == 0 || v > buckets[id].max {
			buckets[id].max = v
		}
		buckets[id].size++
	}

	maxGap := 0
	prev := make([]Pair, 0, 1)
	for i := 0; i < bucketSize; i++ {
		if buckets[i].size == 0 {
			continue
		}
		if len(prev) != 0 {
			maxGap = max(maxGap, buckets[i].min-prev[0].max)
		}
		prev = buckets[i : i+1]
	}

	return maxGap
}

func findMinMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
