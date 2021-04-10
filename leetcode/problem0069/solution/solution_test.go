package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input, output int
	}{
		{4, 2},
		{8, 2},
	}
	for index, test := range tests {
		output := mySqrt(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("index: %v, expected: %v, output: %v", index, test.output, output)
		}
	}
}
