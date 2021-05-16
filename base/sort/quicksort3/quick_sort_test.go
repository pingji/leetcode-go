package quicksort3

import (
	"leetcode-go/common/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	nums := array.GenerateRandomArray(1000000, 1, 1000000)
	// fmt.Printf("before sort, nums: \n%v\n", nums)

	array.TestSort(nums, QuickSort, "Quick Sort 3")
	// fmt.Printf("after sort, nums: \n%v\n", nums)

	assert.True(array.IsSorted(nums))
}
