package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	houses  []int
	heaters []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		// {
		// 	input: Input{
		// 		[]int{1, 2, 3},
		// 		[]int{2},
		// 	},
		// 	output: 1,
		// },
		// {
		// 	input: Input{
		// 		[]int{1, 2, 3, 4},
		// 		[]int{1, 4},
		// 	},
		// 	output: 1,
		// },
		// {
		// 	input: Input{
		// 		[]int{1, 5},
		// 		[]int{2},
		// 	},
		// 	output: 3,
		// },
		{
			input: Input{
				[]int{1, 2, 3},
				[]int{1, 2, 3},
			},
			output: 0,
		},
	}
	for index, test := range tests {
		output := findRadius(test.input.houses, test.input.heaters)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
