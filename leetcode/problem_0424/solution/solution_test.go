package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  string
		k      int
		output int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}
	for index, test := range tests {
		output := characterReplacement(test.input, test.k)
		t.Logf("index: %v, input: %q, t: %v, output %v", index, test.input, test.k, output)
		assert.Equal(test.output, output)
	}
}
