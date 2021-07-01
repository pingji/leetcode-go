package solution

func numRabbits(answers []int) int {
	count := make(map[int]int)
	for _, answer := range answers {
		count[answer]++
	}
	res := 0
	for k, v := range count {
		t := v / (k + 1)
		if v%(k+1) != 0 {
			t++
		}
		res += t * (k + 1)

	}

	return res
}
