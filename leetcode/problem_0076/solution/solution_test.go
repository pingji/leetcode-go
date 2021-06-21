package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  string
		t      string
		output string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"a", "b", ""},
		{"abc", "ac", "abc"},
	}
	for index, test := range tests {
		output := minWindow(test.input, test.t)
		t.Logf("index: %v, input: %q, t: %q, output %q", index, test.input, test.t, output)
		assert.Equal(test.output, output)
	}
}
