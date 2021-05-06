package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  int
		output [][]string
	}{
		{4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
	}
	for index, test := range tests {
		output := solveNQueens(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		if len(output) != len(test.output) {
			t.Errorf("Failed")
		}
	}
}
