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
		output bool
	}{
		{
			input: Input{
				[]int{2, 5, 6, 0, 0, 1, 2},
				0,
			},
			output: true,
		},
		{
			input: Input{
				[]int{2, 5, 6, 0, 0, 1, 2},
				3,
			},
			output: false,
		},
	}
	for index, test := range tests {
		output := search(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
