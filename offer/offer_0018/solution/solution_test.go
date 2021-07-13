package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	head *ListNode
	val  int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output *ListNode
	}{
		{
			input: Input{
				head: linked_list.InitList([]int{4, 5, 1, 9}),
				val:  5,
			},
			output: linked_list.InitList([]int{4, 1, 9}),
		},
		{
			input: Input{
				head: linked_list.InitList([]int{4, 5, 1, 9}),
				val:  1,
			},
			output: linked_list.InitList([]int{4, 5, 9}),
		},
	}
	for index, test := range tests {
		output := deleteNode(test.input.head, test.input.val)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(output.IsEqual(test.output))
	}
}
