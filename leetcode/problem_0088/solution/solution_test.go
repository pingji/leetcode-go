package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{1, 2, 3, 0, 0, 0},
				3,
				[]int{2, 5, 6},
				3,
			},
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			input: Input{
				[]int{1},
				1,
				[]int{},
				0,
			},
			output: []int{1},
		},
		{
			input: Input{
				[]int{0},
				0,
				[]int{1},
				1,
			},
			output: []int{1},
		},
	}
	for index, test := range tests {
		t.Logf("index: %v, input: %v", index, test.input)
		merge(test.input.nums1, test.input.m, test.input.nums2, test.input.n)
		t.Logf("index: %v, output %v", index, test.input.nums1)
		assert.Equal(test.input.nums1, test.output)
	}
}
