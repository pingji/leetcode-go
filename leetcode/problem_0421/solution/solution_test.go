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
				[]int{3, 10, 5, 25, 2, 8},
			},
			output: 28,
		},
		{
			input: Input{
				[]int{0},
			},
			output: 0,
		},
		{
			input: Input{
				[]int{2, 4},
			},
			output: 6,
		},
		{
			input: Input{
				[]int{8, 10, 2},
			},
			output: 10,
		},
		{
			input: Input{
				[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70},
			},
			output: 127,
		},
	}
	for index, test := range tests {
		output := findMaximumXOR(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
