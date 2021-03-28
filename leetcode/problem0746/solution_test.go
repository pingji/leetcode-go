package leetcode

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, test := range tests {
		output := minCostClimbingStairs(test.input)
		t.Logf("input: %v, output %v, expected: %v", test.input, output, test.output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}
