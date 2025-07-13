package solution

import (
	"sort"
)

// Car 结构体表示一辆车，包含位置和到达目标点的时间
type Car struct {
	pos  int     // 车辆当前位置
	time float64 // 车辆到达目标点所需时间
}

// carFleet 计算车队数量
// 思路：从后往前遍历，如果后面的车到达时间小于等于前面的车，则会被前面的车阻挡，形成车队
func carFleet(target int, position []int, speed []int) int {
	// 创建车辆数组，存储每辆车的位置和到达时间
	cars := make([]Car, len(position))
	if len(cars) == 0 {
		return 0
	}

	// 计算每辆车到达目标点的时间：(目标距离 - 当前位置) / 速度
	for i := range position {
		cars[i].pos = position[i]
		cars[i].time = float64(target-position[i]) / float64(speed[i])
	}

	// 按位置从大到小排序（从后往前）
	// 这样我们可以从最接近目标的车辆开始处理
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	// 统计车队数量
	count := 1 // 至少有一个车队（第一辆车）
	for i := 1; i < len(cars); i++ {
		// 如果当前车到达时间小于等于前一辆车，说明会被阻挡
		// 此时当前车的时间应该等于前一辆车的时间（被阻挡）
		if cars[i].time <= cars[i-1].time {
			cars[i].time = cars[i-1].time
		} else {
			// 如果当前车到达时间大于前一辆车，说明不会被阻挡，形成新的车队
			count++
		}
	}
	return count
}
