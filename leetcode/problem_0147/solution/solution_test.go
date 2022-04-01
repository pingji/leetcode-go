package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output []int
	}{
		{
			input: Input{
				[]int{4, 2, 1, 3},
			},
			output: []int{1, 2, 3, 4},
		},
		{
			input: Input{
				[]int{-1, 5, 3, 4, 0},
			},
			output: []int{-1, 0, 3, 4, 5},
		},
	}
	for index, test := range tests {
		input := linked_list.InitList(test.input.nums)
		outputA := linked_list.InitList(test.output)
		t.Logf("index: %v, input: %v", index, test.input)
		output := insertionSortList(input)
		t.Logf("index: %v, output: %v", index, output)
		assert.True(output.IsEqual(outputA))
	}
}
