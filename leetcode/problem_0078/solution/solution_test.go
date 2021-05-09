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
			{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
		}},
		{[]int{0}, [][]int{
			{}, {0},
		}},
	}
	assert := assert.New(t)
	for index, test := range tests {
		output := subsets(test.input)
		t.Logf("index: %v, input: %v, output: %v", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
