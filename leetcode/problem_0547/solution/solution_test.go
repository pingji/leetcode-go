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
			{1, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		}, 2},
		{[][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		}, 3},
		{[][]int{
			{1, 0, 0, 1},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{1, 0, 1, 1},
		}, 1},
	}
	for index, test := range tests {
		output := findCircleNum(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
