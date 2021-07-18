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
				5,
			},
			output: 5,
		},
		{
			input: Input{
				45,
			},
			output: 134903163,
		},
		{
			input: Input{
				95,
			},
			output: 407059028,
		},
	}
	for index, test := range tests {
		output := fib(test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
