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
				[]int{1, 2, 3},
			},
			output: 0,
		},
		{
			input: Input{
				[]int{1, 2, 3},
			},
			output: 0,
		},
		{
			input: Input{
				[]int{1, 2, 3},
			},
			output: 0,
		},
	}
	for index, test := range tests {
		output := foo(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
