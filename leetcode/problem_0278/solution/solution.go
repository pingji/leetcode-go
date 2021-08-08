package solution

var bad = 1

func isBadVersion(version int) bool {
	return version >= bad
}

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
