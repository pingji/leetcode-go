package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s string
	k int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output string
	}{
		{
			input: Input{
				"abcdefg",
				2,
			},
			output: "bacdfeg",
		},
		{
			input: Input{
				"abcd",
				2,
			},
			output: "bacd",
		},
	}
	for index, test := range tests {
		output := reverseStr(test.input.s, test.input.k)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
