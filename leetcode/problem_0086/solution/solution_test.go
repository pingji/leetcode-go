package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	head *ListNode
	x    int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output *ListNode
	}{
		{
			input: Input{
				head: linked_list.InitList([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			output: linked_list.InitList([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			input: Input{
				head: linked_list.InitList([]int{2, 1}),
				x:    2,
			},
			output: linked_list.InitList([]int{1, 2}),
		},
	}
	for index, test := range tests {
		output := partition(test.input.head, test.input.x)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(output.IsEqual(test.output))
	}
}
