package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	l1 *ListNode
	l2 *ListNode
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output *ListNode
	}{
		{
			input: Input{
				l1: linked_list.InitList([]int{1, 2, 4}),
				l2: linked_list.InitList([]int{1, 3, 4}),
			},
			output: linked_list.InitList([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for index, test := range tests {
		output := mergeTwoLists(test.input.l1, test.input.l2)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(output.IsEqual(test.output))
	}
}
