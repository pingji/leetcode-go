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
				4,
			},
			output: 4,
		},
		{
			input: Input{
				25,
			},
			output: 1389537,
		},
		{
			input: Input{
				0,
			},
			output: 0,
		},
	}
	for index, test := range tests {
		output := tribonacci(test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
