package solution

func numberOfSubarrays(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	oddnum := 0
	for i := 0; i < len(nums); i++ {
		oddnum += nums[i] & 1
		if hash[oddnum-k] > 0 {
			count += hash[oddnum-k]
		}
		hash[oddnum]++
	}
	return count
}
