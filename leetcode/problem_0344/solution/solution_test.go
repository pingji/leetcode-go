package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s []byte
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []byte
	}{
		{
			input: Input{
				[]byte{'h', 'e', 'l', 'l', 'o'},
			},
			output: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			input: Input{
				[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			output: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for index, test := range tests {
		input := make([]byte, len(test.input.s))
		copy(input, test.input.s)

		reverseString(test.input.s)

		output := make([]byte, len(test.input.s))
		copy(output, test.input.s)

		t.Logf("index: %v, input: %v, output %v", index, input, output)
		assert.Equal(test.output, output)
	}
}
