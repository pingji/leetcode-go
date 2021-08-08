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
				5,
			},
			output: 4,
		},
		{
			input: Input{
				1,
			},
			output: 1,
		},
	}
	for index, test := range tests {
		bad = test.output
		output := firstBadVersion(test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
