package solution

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stations := map[int][]int{}
	for i, route := range routes {
		for _, station := range route {
			stations[station] = append(stations[station], i)
		}
	}
	visited := make([]bool, len(routes))
	q := []int{source}
	buses := 0

	for len(q) != 0 {
		size := len(q)
		buses++
		for i := 0; i < size; i++ {
			curStation := q[0]
			q = q[1:]

			for _, route := range stations[curStation] {
				if visited[route] {
					continue
				}
				visited[route] = true
				for _, station := range routes[route] {
					if station == target {
						return buses
					}
					q = append(q, station)
				}
			}
		}
	}

	return -1
}
