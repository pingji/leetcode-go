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
		output []int
	}{
		{
			input: Input{
				[]int{5, 2, 3, 1},
			},
			output: []int{1, 2, 3, 5},
		},
		{
			input: Input{
				[]int{5, 1, 1, 2, 0, 0},
			},
			output: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for index, test := range tests {
		output := sortArray(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
