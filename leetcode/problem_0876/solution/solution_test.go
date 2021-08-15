package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	head *ListNode
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				head: linked_list.InitList([]int{1, 2, 3, 4, 5}),
			},
			output: 3,
		},
		{
			input: Input{
				head: linked_list.InitList([]int{1, 2, 3, 4, 5, 6}),
			},
			output: 4,
		},
	}
	for index, test := range tests {
		output := middleNode(test.input.head)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output.Val)
	}
}
