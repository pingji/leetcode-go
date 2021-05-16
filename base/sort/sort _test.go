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
	nums1 := array.GenerateRandomArray(1000000, 1, 100)
	// nums1 := array.GenerateRandomArray(1000000, 1, 1000000)
	nums2 := make([]int, len(nums1))
	copy(nums2, nums1)
	nums3 := make([]int, len(nums1))
	copy(nums3, nums1)

	array.TestSort(nums1, quicksort1.QuickSort, "Quick Sort 1")
	array.TestSort(nums2, quicksort2.QuickSort, "Quick Sort 2")
	array.TestSort(nums3, quicksort3.QuickSort, "Quick Sort 3")

	assert.True(array.IsSorted(nums1))
	assert.True(array.IsSorted(nums2))
	assert.True(array.IsSorted(nums3))
}
