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
				[]int{1, 2, 3, 4},
			},
			output: []int{1, 4, 2, 3},
		},
		// {
		// 	input: Input{
		// 		[]int{1, 2, 3, 4, 5},
		// 	},
		// 	output: []int{1, 5, 2, 4, 3},
		// },
	}
	for index, test := range tests {
		input := linked_list.InitList(test.input.nums)
		output := linked_list.InitList(test.output)
		t.Logf("index: %v, input: %v", index, test.input)
		reorderList(input)
		t.Logf("index: %v, output: %v", index, test.output)
		assert.True(output.IsEqual(input))
	}
}
