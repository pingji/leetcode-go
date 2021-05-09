package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert := assert.New(t)
	for index, test := range tests {
		output := solveNQueens(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
