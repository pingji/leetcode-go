package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  *ListNode
		output *ListNode
	}{
		{
			input:  linked_list.InitList([]int{1, 2, 3, 4}),
			output: linked_list.InitList([]int{2, 1, 4, 3}),
		},
		{
			input:  linked_list.InitList([]int{}),
			output: linked_list.InitList([]int{}),
		},
		{
			input:  linked_list.InitList([]int{1}),
			output: linked_list.InitList([]int{1}),
		},
	}
	for index, test := range tests {
		output := swapPairs(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(output.IsEqual(test.output))
	}
}
