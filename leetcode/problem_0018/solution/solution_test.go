package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums   []int
	target int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output [][]int
	}{
		{
			input: Input{
				[]int{1, 0, -1, 0, -2, 2},
				0,
			},
			output: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			input: Input{
				[]int{2, 2, 2, 2, 2},
				8,
			},
			output: [][]int{{2, 2, 2, 2}},
		},
		{
			input: Input{
				[]int{1, -2, -5, -4, -3, 3, 3, 5},
				-11,
			},
			output: [][]int{{-5, -4, -3, 1}},
		},
	}
	for index, test := range tests {
		output := fourSum(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
