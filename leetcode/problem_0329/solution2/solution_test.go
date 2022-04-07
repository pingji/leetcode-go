package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	matrix [][]int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}},
			},
			output: 4,
		},
		{
			input: Input{
				[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}},
			},
			output: 4,
		},
		{
			input: Input{
				[][]int{{1}},
			},
			output: 1,
		},
	}
	for index, test := range tests {
		output := longestIncreasingPath(test.input.matrix)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
