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
		// {
		// 	input: Input{
		// 		[]int{-1,2,1,-4},
		// 		1,
		// 	},
		// 	output: 2,
		// },
		// {
		// 	input: Input{
		// 		[]int{0,0,0},
		// 		1,
		// 	},
		// 	output: 0,
		// },
		{
			input: Input{
				[]int{1, 1, 1, 1},
				0,
			},
			output: 3,
		},
	}
	for index, test := range tests {
		output := threeSumClosest(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
