package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input1, input2 []int
		output         []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
	}
	for index, test := range tests {
		list1 := linked_list.InitList(test.input1)
		list2 := linked_list.InitList(test.input2)
		output := mergeTwoLists(list1, list2)
		t.Logf("index: %v, input1: %v, input2: %v, output %v", index, list1.ToString(), list2.ToString(), output.ToString())
		assert.True(linked_list.InitList(test.output).IsEqual(output))
	}
}
