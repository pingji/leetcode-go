package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}
	assert := assert.New(t)
	for index, test := range tests {
		output := permute(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
