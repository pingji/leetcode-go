package solution

import (
	"testing"
)

type Input struct {
	g []int
	s []int
}

func TestProblem(t *testing.T) {
	tests := []struct {
		input  Input
		output int
	}{
		{
			Input{
				[]int{1, 2, 3},
				[]int{1, 1},
			},
			1,
		},
		{
			Input{
				[]int{1, 2},
				[]int{1, 2, 3},
			},
			2,
		},
	}
	for index, test := range tests {
		output := findContentChildren(test.input.g, test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("Error! expected: %v, output: %v", test.output, output)
		}
	}
}
