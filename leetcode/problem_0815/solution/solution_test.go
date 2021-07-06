package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	routes [][]int
	source int
	target int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			Input{
				routes: [][]int{{1, 2, 7}, {3, 6, 7}},
				source: 1,
				target: 6,
			},
			2,
		},
		{
			Input{
				routes: [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}},
				source: 15,
				target: 12,
			},
			-1,
		},
	}
	for index, test := range tests {
		output := numBusesToDestination(test.input.routes, test.input.source, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
