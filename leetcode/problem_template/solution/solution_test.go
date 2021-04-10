package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input, output int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	for index, test := range tests {
		output := foo(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("index: %v, expected: %v, output: %v", index, test.output, output)
		}
	}
}
