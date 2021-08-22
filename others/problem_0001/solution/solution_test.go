package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	arr [][]int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			Input{
				[][]int{
					{1, 3, 5, 7},
					{2, 4, 6, 8},
					{0, 9, 10, 11},
				},
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			Input{
				[][]int{
					{1, 5, 6, 8},
					{2, 4, 10, 12},
					{3, 7, 9, 11},
					{13, 14, 15, 16},
				},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}
	for index, test := range tests {
		output := mergeKArrays(test.input.arr)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
