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
		output bool
	}{
		{
			input: Input{
				2,
				[][]int{{1, 0}},
			},
			output: true,
		},
		{
			input: Input{
				2,
				[][]int{{1, 0}, {0, 1}},
			},
			output: false,
		},
	}
	for index, test := range tests {
		output := canFinish(test.input.numCourses, test.input.prerequisites)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
