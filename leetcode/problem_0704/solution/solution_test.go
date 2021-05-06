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
		{Input{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{Input{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
	}
	for index, test := range tests {
		output := search(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("index: %v, expected: %v, output: %v", index, test.output, output)
		}
	}
}
