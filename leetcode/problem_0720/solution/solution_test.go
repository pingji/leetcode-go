package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	words []string
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output string
	}{
		{
			input: Input{
				[]string{"w", "wo", "wor", "worl", "world"},
			},
			output: "world",
		},
		{
			input: Input{
				[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			},
			output: "apple",
		},
	}
	for index, test := range tests {
		output := longestWord(test.input.words)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
