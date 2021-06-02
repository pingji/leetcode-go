package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums  []int
	lower int
	upper int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				nums:  []int{-2, 5, -1},
				lower: -2,
				upper: 2,
			},
			output: 3,
		},
		{
			input: Input{
				nums:  []int{0},
				lower: 0,
				upper: 0,
			},
			output: 1,
		},
	}
	for index, test := range tests {
		output := countRangeSum(test.input.nums, test.input.lower, test.input.upper)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
