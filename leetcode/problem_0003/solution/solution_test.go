package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}
	for index, test := range tests {
		output := lengthOfLongestSubstring(test.input)
		t.Logf("index: %v, input: %q, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
