package array

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomArray(size int, rangeL int, rangeR int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(rangeR-rangeL+1) + rangeL
	}
	return nums
}

func TestSort(nums []int, sort func([]int), sortName string) {
	begin := time.Now()
	sort(nums)
	end := time.Now()
	duration := end.Sub(begin)
	fmt.Printf("%s, time used: %v\n", sortName, duration)
}

func IsSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
