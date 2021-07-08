package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums   []int
	target int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			Input{
				[]int{2, 7, 11, 15},
				9,
			},
			[]int{0, 1},
		},
		{
			Input{
				[]int{3, 2, 4},
				6,
			},
			[]int{1, 2},
		},
		{
			Input{
				[]int{3, 3},
				6,
			},
			[]int{0, 1},
		},
	}
	for index, test := range tests {
		output := twoSum(test.input.nums, test.input.target)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
