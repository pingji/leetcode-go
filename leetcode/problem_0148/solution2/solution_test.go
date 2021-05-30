package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
		{[]int{}, []int{}},
	}
	for index, test := range tests {
		head := linked_list.InitList(test.input)
		output := sortList(head)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.True(linked_list.InitList(test.output).IsEqual(output))
	}
}
