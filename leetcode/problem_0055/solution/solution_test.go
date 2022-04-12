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
				[]int{2, 3, 1, 1, 4},
			},
			output: true,
		},
		{
			input: Input{
				[]int{3, 2, 1, 0, 4},
			},
			output: false,
		},
	}
	for index, test := range tests {
		output := canJump(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
