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
		output string
	}{
		{
			input: Input{
				"abcdabcaba",
			},
			output: "b",
		},
		{
			input: Input{
				"abcdabcabab",
			},
			output: "c",
		},
	}
	for index, test := range tests {
		output := secondFreqChar(test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
