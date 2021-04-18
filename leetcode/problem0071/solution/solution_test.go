package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
	}
	for index, test := range tests {
		output := simplifyPath(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("Error! expected: %v, output: %v", test.output, output)
		}
	}
}
