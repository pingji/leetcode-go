package solution

// findOrder 返回完成所有课程的学习顺序
// 如果无法完成所有课程，返回空数组
// 使用拓扑排序算法，如果有多个有效顺序，返回任意一个
func findOrder(numCourses int, prerequisites [][]int) []int {
	// edges[i] 表示从课程i出发可以到达的所有课程（邻接表）
	edges := make([][]int, numCourses)
	// indeg[i] 表示课程i的入度（有多少课程指向它）
	indeg := make([]int, numCourses)
	// result 存储拓扑排序的结果，即学习顺序
	result := []int{}
	// empty 空数组，当无法完成所有课程时返回
	empty := []int{}

	// 构建有向图和计算入度
	for _, info := range prerequisites {
		// info[1] -> info[0] 表示必须先修info[1]才能修info[0]
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++ // info[0]的入度+1
	}

	// 将所有入度为0的课程加入队列（这些课程可以立即学习）
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	// 拓扑排序过程
	for len(q) > 0 {
		u := q[0] // 取出队首课程
		q = q[1:]
		result = append(result, u) // 将课程u加入学习顺序

		// 遍历所有以u为前置条件的课程
		for _, v := range edges[u] {
			indeg[v]--         // 课程v的入度减1
			if indeg[v] == 0 { // 如果课程v的入度变为0，可以学习
				q = append(q, v)
			}
		}
	}

	// 判断是否完成了所有课程
	// 如果拓扑排序结果的长度等于课程总数，说明没有环，可以完成所有课程
	if len(result) == numCourses {
		return result
	}
	// 否则说明存在环，无法完成所有课程，返回空数组
	return empty
}
