package solution

import (
	"testing"
)

type Input struct {
	nums   []int
	target int
}

func TestProblem(t *testing.T) {
	tests := []struct {
		input  Input
		output int
	}{
		{Input{[]int{1, 3, 5, 6}, 5}, 2},
		{Input{[]int{1, 3, 5, 6}, 2}, 1},
		{Input{[]int{1, 3, 5, 6}, 7}, 4},
		{Input{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for index, test := range tests {
		output := searchInsert(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("index: %v, expected: %v, output: %v", index, test.output, output)
		}
	}
}
