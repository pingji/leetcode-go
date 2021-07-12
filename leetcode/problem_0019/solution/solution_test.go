package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	head *ListNode
	n    int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output *ListNode
	}{
		{
			input: Input{
				head: linked_list.InitList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			output: linked_list.InitList([]int{1, 2, 3, 5}),
		},
		{
			input: Input{
				head: linked_list.InitList([]int{1}),
				n:    1,
			},
			output: linked_list.InitList([]int{}),
		},
		{
			input: Input{
				head: linked_list.InitList([]int{1, 2}),
				n:    1,
			},
			output: linked_list.InitList([]int{1}),
		},
	}
	for index, test := range tests {
		output := removeNthFromEnd(test.input.head, test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(output.IsEqual(test.output))
	}
}
