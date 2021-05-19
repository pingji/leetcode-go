package selectsort

import (
	"fmt"
	"leetcode-go/common/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	nums := array.GenerateRandomArray(10, 1, 10)
	fmt.Printf("before sort, nums: \n%v\n", nums)

	array.TestSort(nums, SelectSort, "Select Sort")
	fmt.Printf("after sort, nums: \n%v\n", nums)

	assert.True(array.IsSorted(nums))
}
