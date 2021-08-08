package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
	k    int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{1, 1, 1, 2, 2, 3},
				2,
			},
			output: []int{1, 2},
		},
		{
			input: Input{
				[]int{1},
				1,
			},
			output: []int{1},
		},
	}
	for index, test := range tests {
		output := topKFrequent(test.input.nums, test.input.k)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
