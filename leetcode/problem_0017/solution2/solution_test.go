package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}
	assert := assert.New(t)
	for index, test := range tests {
		output := letterCombinations(test.input)
		t.Logf("index: %v, input: %q, output: %q", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
