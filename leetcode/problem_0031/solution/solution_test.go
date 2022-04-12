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
				[]int{1, 2, 3},
			},
			output: []int{1, 3, 2},
		},
		{
			input: Input{
				[]int{3, 2, 1},
			},
			output: []int{1, 2, 3},
		},
		{
			input: Input{
				[]int{1, 1, 5},
			},
			output: []int{1, 5, 1},
		},
		{
			input: Input{
				[]int{1, 5, 1},
			},
			output: []int{5, 1, 1},
		},
	}
	for index, test := range tests {
		t.Logf("index: %v, input: %v", index, test.input)
		nextPermutation(test.input.nums)
		t.Logf("index: %v, output %v", index, test.input)
		assert.Equal(test.input.nums, test.output)
	}
}
