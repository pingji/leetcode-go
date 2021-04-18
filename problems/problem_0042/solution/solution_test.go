package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for index, test := range tests {
		output := trap(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("Error! expected: %v, output: %v", test.output, output)
		}
	}
}
