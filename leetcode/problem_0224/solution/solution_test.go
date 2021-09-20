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
				"1 + 1",
			},
			output: 2,
		},
		{
			input: Input{
				" 2-1 + 2 ",
			},
			output: 3,
		},
		{
			input: Input{
				"(1+(4+5+2)-3)+(6+8)",
			},
			output: 23,
		},
	}
	for index, test := range tests {
		output := calculate(test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
