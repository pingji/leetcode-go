package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	head *ListNode
	k    int
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
				k:    2,
			},
			output: linked_list.InitList([]int{4, 5, 1, 2, 3}),
		},
		{
			input: Input{
				head: linked_list.InitList([]int{0, 1, 2}),
				k:    4,
			},
			output: linked_list.InitList([]int{2, 0, 1}),
		},
	}
	for index, test := range tests {
		output := rotateRight(test.input.head, test.input.k)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(output.IsEqual(test.output))
	}
}
