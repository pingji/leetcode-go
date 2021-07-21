package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	numbers []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				[]int{3, 4, 5, 1, 2},
			},
			output: 1,
		},
		{
			input: Input{
				[]int{2, 2, 2, 0, 1},
			},
			output: 0,
		},
	}
	for index, test := range tests {
		output := minArray(test.input.numbers)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
