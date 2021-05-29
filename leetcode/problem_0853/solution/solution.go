package solution

import (
	"sort"
)

type Car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	if len(cars) == 0 {
		return 0
	}
	for i, _ := range position {
		cars[i].pos = position[i]
		cars[i].time = float64(target-position[i]) / float64(speed[i])
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	count := 1
	for i := 1; i < len(cars); i++ {
		if cars[i].time <= cars[i-1].time {
			cars[i].time = cars[i-1].time
		} else {
			count++
		}
	}
	return count
}
