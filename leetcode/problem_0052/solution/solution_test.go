package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{4, 2}, {1, 1},
	}
	assert := assert.New(t)
	for index, test := range tests {
		output := totalNQueens(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
