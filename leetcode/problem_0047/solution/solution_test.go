package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{[]int{1, 2, 1}, [][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		}},
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}
	for index, test := range tests {
		output := permuteUnique(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		// t.Error("Failed")
	}
}
