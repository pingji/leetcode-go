package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{[]int{4, 2}, [][]int{
			{2, 4},
			{3, 4},
			{2, 3},
			{1, 2},
			{1, 3},
			{1, 4},
		}},
	}
	for index, test := range tests {
		output := combine(test.input[0], test.input[1])
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		t.Error("Fail")
	}
}
