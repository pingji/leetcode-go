package solution

import (
	"leetcode-go/common/array"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
	}
	for _, test := range tests {
		output := nextGreaterElements(test.nums)
		t.Logf("input: %v, output %v, expected: %v", test.nums, output, test.output)
		if !array.IsEqual(output, test.output) {
			t.Errorf("Failed")
		}
	}
}
