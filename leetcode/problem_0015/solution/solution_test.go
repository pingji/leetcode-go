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
		output [][]int
	}{
		{
			input: Input{
				[]int{-1, 0, 1, 2, -1, -4},
			},
			output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input: Input{
				[]int{},
			},
			output: [][]int{},
		},
		{
			input: Input{
				[]int{0},
			},
			output: [][]int{},
		},
	}
	for index, test := range tests {
		output := threeSum(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
