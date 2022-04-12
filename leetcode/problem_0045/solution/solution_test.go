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
				[]int{2, 3, 1, 1, 4},
			},
			output: 2,
		},
		{
			input: Input{
				[]int{2, 3, 0, 1, 4},
			},
			output: 2,
		},
		{
			input: Input{
				[]int{2, 1},
			},
			output: 1,
		},
	}
	for index, test := range tests {
		output := jump(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
