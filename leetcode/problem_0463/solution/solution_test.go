package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  [][]int
		output int
	}{
		{[][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		}, 16},
		{[][]int{
			{1},
		}, 4},
		{[][]int{
			{1, 0},
		}, 4},
	}
	for index, test := range tests {
		output := islandPerimeter(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
