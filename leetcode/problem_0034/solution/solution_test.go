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
		output []int
	}{
		{
			input: Input{
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			output: []int{3, 4},
		},
		{
			input: Input{
				[]int{5, 7, 7, 8, 8, 10},
				6,
			},
			output: []int{-1, -1},
		},
		{
			input: Input{
				[]int{},
				0,
			},
			output: []int{-1, -1},
		},
	}
	for index, test := range tests {
		output := searchRange(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
