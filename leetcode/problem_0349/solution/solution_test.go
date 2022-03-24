package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums1 []int
	nums2 []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{1, 2, 2, 1},
				[]int{2, 2},
			},
			output: []int{2},
		},
		{
			input: Input{
				[]int{4, 9, 5},
				[]int{9, 4, 9, 8, 4},
			},
			output: []int{4, 9},
		},
	}
	for index, test := range tests {
		output := intersection(test.input.nums1, test.input.nums2)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
