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
				"tree",
			},
			output: "eetr",
		},
		{
			input: Input{
				"cccaaa",
			},
			output: "cccaaa",
		},
		{
			input: Input{
				"Aabb",
			},
			output: "bbAa",
		},
	}
	for index, test := range tests {
		output := frequencySort(test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
