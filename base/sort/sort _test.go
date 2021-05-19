package sort

import (
	"leetcode-go/base/sort/insertsort"
	"leetcode-go/base/sort/mergesort1"
	"leetcode-go/base/sort/mergesort2"
	"leetcode-go/base/sort/quicksort1"
	"leetcode-go/base/sort/quicksort2"
	"leetcode-go/base/sort/quicksort3"
	"leetcode-go/base/sort/selectsort"
	"leetcode-go/common/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert := assert.New(t)
	// nums1 := array.GenerateRandomArray(100000, 1, 100)
	nums1 := array.GenerateRandomArray(100000, 1, 100000)
	nums2 := make([]int, len(nums1))
	copy(nums2, nums1)
	nums3 := make([]int, len(nums1))
	copy(nums3, nums1)
	nums4 := make([]int, len(nums1))
	copy(nums4, nums1)
	nums5 := make([]int, len(nums1))
	copy(nums5, nums1)
	nums6 := make([]int, len(nums1))
	copy(nums6, nums1)
	nums7 := make([]int, len(nums1))
	copy(nums7, nums1)

	array.TestSort(nums1, selectsort.SelectSort, "Select Sort")
	array.TestSort(nums2, insertsort.InsertSort, "Insert Sort")
	array.TestSort(nums3, mergesort1.MergeSort, "Merge Sort 1")
	array.TestSort(nums4, mergesort2.MergeSort, "Merge Sort 2")
	array.TestSort(nums5, quicksort1.QuickSort, "Quick Sort 1")
	array.TestSort(nums6, quicksort2.QuickSort, "Quick Sort 2")
	array.TestSort(nums7, quicksort3.QuickSort, "Quick Sort 3")

	assert.True(array.IsSorted(nums1))
	assert.True(array.IsSorted(nums2))
	assert.True(array.IsSorted(nums3))
	assert.True(array.IsSorted(nums4))
	assert.True(array.IsSorted(nums5))
	assert.True(array.IsSorted(nums6))
	assert.True(array.IsSorted(nums7))
}
