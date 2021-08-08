package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{-4, -1, 0, 3, 10},
			},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			input: Input{
				[]int{-7, -3, 2, 3, 11},
			},
			output: []int{4, 9, 9, 49, 121},
		},
	}
	for index, test := range tests {
		output := sortedSquares(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
