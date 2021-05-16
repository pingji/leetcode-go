package sort

import (
	"leetcode-go/base/sort/quicksort1"
	"leetcode-go/base/sort/quicksort2"
	"leetcode-go/base/sort/quicksort3"
	"leetcode-go/common/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort1(t *testing.T) {
	assert := assert.New(t)

	nums := array.GenerateRandomArray(1000000, 1, 1000000)
	// fmt.Printf("before sort, nums: \n%v\n", nums)

	array.TestSort(nums, quicksort1.QuickSort, "Quick Sort 1")
	// fmt.Printf("after sort, nums: \n%v\n", nums)

	assert.True(array.IsSorted(nums))
}

func TestQuickSort2(t *testing.T) {
	assert := assert.New(t)

	nums := array.GenerateRandomArray(1000000, 1, 1000000)
	// fmt.Printf("before sort, nums: \n%v\n", nums)

	array.TestSort(nums, quicksort2.QuickSort, "Quick Sort 2")
	// fmt.Printf("after sort, nums: \n%v\n", nums)

	assert.True(array.IsSorted(nums))
}

func TestQuickSort3(t *testing.T) {
	assert := assert.New(t)

	nums := array.GenerateRandomArray(1000000, 1, 1000000)
	// fmt.Printf("before sort, nums: \n%v\n", nums)

	array.TestSort(nums, quicksort3.QuickSort, "Quick Sort 3")
	// fmt.Printf("after sort, nums: \n%v\n", nums)

	assert.True(array.IsSorted(nums))
}
