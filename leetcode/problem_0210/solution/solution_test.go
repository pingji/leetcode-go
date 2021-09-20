package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	numCourses    int
	prerequisites [][]int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				2,
				[][]int{{1, 0}},
			},
			output: []int{0, 1},
		},
		{
			input: Input{
				4,
				[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			output: []int{0, 1, 2, 3},
		},
	}
	for index, test := range tests {
		output := findOrder(test.input.numCourses, test.input.prerequisites)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
