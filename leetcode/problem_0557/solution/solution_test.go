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
				"Let's take LeetCode contest",
			},
			output: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for index, test := range tests {
		output := reverseWords(test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
