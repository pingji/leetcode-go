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
			output: 1,
		},
		{
			input: Input{
				3,
			},
			output: 2,
		},
		{
			input: Input{
				4,
			},
			output: 3,
		},
	}
	for index, test := range tests {
		output := fib(test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
