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
		output float64
	}{
		{
			input: Input{
				[]int{1, 3},
				[]int{2},
			},
			output: 2.00000,
		},
		{
			input: Input{
				[]int{1, 2},
				[]int{3, 4},
			},
			output: 2.50000,
		},
		{
			input: Input{
				[]int{0, 0},
				[]int{0, 0},
			},
			output: 0.00000,
		},
		{
			input: Input{
				[]int{},
				[]int{1},
			},
			output: 1.00000,
		},
		{
			input: Input{
				[]int{2},
				[]int{},
			},
			output: 2.00000,
		},
	}
	for index, test := range tests {
		output := findMedianSortedArrays(test.input.nums1, test.input.nums2)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
