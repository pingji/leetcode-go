package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums [][]int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			input: Input{
				[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			},
			output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for index, test := range tests {
		output := spiralOrder(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
