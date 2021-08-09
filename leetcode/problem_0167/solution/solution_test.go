package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	numbers []int
	target  int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{2, 7, 11, 15},
				9,
			},
			output: []int{1, 2},
		},
		{
			input: Input{
				[]int{2, 3, 4},
				6,
			},
			output: []int{1, 3},
		},
		{
			input: Input{
				[]int{-1, 0},
				-1,
			},
			output: []int{1, 2},
		},
	}
	for index, test := range tests {
		output := twoSum(test.input.numbers, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
