package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s string
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				"42",
			},
			output: 42,
		},
		{
			input: Input{
				"   -42",
			},
			output: -42,
		},
		{
			input: Input{
				"4193 with words",
			},
			output: 4193,
		},
	}
	for index, test := range tests {
		output := myAtoi(test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
