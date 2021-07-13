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
		output []int
	}{
		{
			input:  linked_list.InitList([]int{1, 3, 2}),
			output: []int{2, 3, 1},
		},
	}
	for index, test := range tests {
		output := reversePrint(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
