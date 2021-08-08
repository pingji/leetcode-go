package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
	k    int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{1, 2, 3, 4, 5, 6, 7},
				3,
			},
			output: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			input: Input{
				[]int{-1, -100, 3, 99},
				2,
			},
			output: []int{3, 99, -1, -100},
		},
	}
	for index, test := range tests {
		t.Logf("index: %v, input: %v", index, test.input)
		rotate(test.input.nums, test.input.k)
		t.Logf("index: %v, output %v", index, test.input)
		assert.Equal(test.output, test.input.nums)
	}
}
