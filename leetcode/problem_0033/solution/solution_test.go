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
		output int
	}{
		{
			input: Input{
				[]int{4, 5, 6, 7, 0, 1, 2},
				0,
			},
			output: 4,
		},
		{
			input: Input{
				[]int{4, 5, 6, 7, 0, 1, 2},
				3,
			},
			output: -1,
		},
		{
			input: Input{
				[]int{1},
				0,
			},
			output: -1,
		},
	}
	for index, test := range tests {
		output := search(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
