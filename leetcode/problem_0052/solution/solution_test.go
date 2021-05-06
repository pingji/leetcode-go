package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{4, 2}, {1, 1},
	}
	for index, test := range tests {
		output := totalNQueens(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}
