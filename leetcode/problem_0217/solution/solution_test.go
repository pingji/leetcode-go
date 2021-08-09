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
		output bool
	}{
		{
			input: Input{
				[]int{1, 2, 3, 1},
			},
			output: true,
		},
		{
			input: Input{
				[]int{1, 2, 3, 4},
			},
			output: false,
		},
		{
			input: Input{
				[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			output: true,
		},
	}
	for index, test := range tests {
		output := containsDuplicate(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
