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
		// {
		// 	input: Input{
		// 		[]int{1, 3, 5, 7, 2, 4, 6, 8},
		// 		4,
		// 	},
		// 	output: []int{1, 2, 3, 4},
		// },
		{
			input: Input{
				[]int{1, 2, 3},
				0,
			},
			output: []int{},
		},
	}
	for index, test := range tests {
		output := smallestK(test.input.nums, test.input.k)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
