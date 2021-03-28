package leetcode

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{2, 7, 9, 3, 1}, 22},
		{[]int{}, 0},
		{[]int{-1, -2}, -1},
	}
	for _, test := range tests {
		output := maxSubArray(test.input)
		t.Logf("input: %v, output %v, expected: %v", test.input, output, test.output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}
