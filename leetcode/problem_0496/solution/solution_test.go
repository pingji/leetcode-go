package solution

import (
	"leetcode-go/common/array"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		output []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}
	for _, test := range tests {
		output := nextGreaterElement(test.nums1, test.nums2)
		t.Logf("input: %v, %v, output %v, expected: %v", test.nums1, test.nums2, output, test.output)
		if !array.IsEqual(output, test.output) {
			t.Errorf("Failed")
		}
	}
}
