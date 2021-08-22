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
				"3+2*2",
			},
			output: 7,
		},
		// {
		// 	input: Input{
		// 		" 3/2 ",
		// 	},
		// 	output: 1,
		// },
		// {
		// 	input: Input{
		// 		" 3+5 / 2 ",
		// 	},
		// 	output: 5,
		// },
	}
	for index, test := range tests {
		output := calculate(test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
