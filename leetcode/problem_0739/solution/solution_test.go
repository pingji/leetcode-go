package solution

import (
	"leetcode-go/common/array"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, test := range tests {
		output := dailyTemperatures(test.input)
		t.Logf("input: %v, output %v, expected: %v", test.input, output, test.output)
		if !array.IsEqual(output, test.output) {
			t.Errorf("Failed")
		}
	}
}
