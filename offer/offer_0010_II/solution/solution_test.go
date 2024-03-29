package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	n int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				2,
			},
			output: 2,
		},
		{
			input: Input{
				3,
			},
			output: 3,
		},
		{
			input: Input{
				7,
			},
			output: 21,
		},
		{
			input: Input{
				0,
			},
			output: 1,
		},
	}
	for index, test := range tests {
		output := numWays(test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
