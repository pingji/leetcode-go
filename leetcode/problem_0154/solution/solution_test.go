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
				[]int{1, 3, 5},
			},
			output: 1,
		},
		{
			input: Input{
				[]int{2, 2, 2, 0, 1},
			},
			output: 0,
		},
		{
			input: Input{
				[]int{1, 3, 1},
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
