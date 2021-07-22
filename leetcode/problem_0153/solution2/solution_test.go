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
		output int
	}{
		{
			input: Input{
				[]int{3, 4, 5, 1, 2},
			},
			output: 1,
		},
		{
			input: Input{
				[]int{4, 5, 6, 7, 0, 1, 2},
			},
			output: 0,
		},
		{
			input: Input{
				[]int{11, 13, 15, 17},
			},
			output: 11,
		},
		{
			input: Input{
				[]int{1},
			},
			output: 1,
		},
		{
			input: Input{
				[]int{2, 1},
			},
			output: 1,
		},
	}
	for index, test := range tests {
		output := findMin(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
