package solution

func merge(sum []int, lower int, upper int, left int, right int) int {
	if left == right {
		return 0
	}
	mid := (right + left) / 2
	n1 := merge(sum, lower, upper, left, mid)
	n2 := merge(sum, lower, upper, mid+1, right)
	count := n1 + n2

	l := mid + 1
	r := mid + 1

	for i := left; i <= mid; i++ {
		for l <= right && sum[l]-sum[i] < lower {
			l++
		}
		for r <= right && sum[r]-sum[i] <= upper {
			r++
		}
		count += (r - l)
	}

	// merge sort
	temp := make([]int, right-left+1)
	copy(temp, sum[left:right+1])

	tmpMid := mid - left
	tmpRight := right - left
	p1 := 0
	p2 := mid + 1 - left
	for k := left; k <= right; k++ {
		if p1 > tmpMid {
			sum[k] = temp[p2]
			p2++
		} else if p2 > tmpRight {
			sum[k] = temp[p1]
			p1++
		} else if temp[p1] < temp[p2] {
			sum[k] = temp[p1]
			p1++
		} else {
			sum[k] = temp[p2]
			p2++
		}
	}
	return count
}

func countRangeSum(nums []int, lower int, upper int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	return merge(sum, lower, upper, 0, len(sum)-1)
}
