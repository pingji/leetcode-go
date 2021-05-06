package solution

import (
	"leetcode-go/common/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  [][]int
		output []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}}, []int{}},
	}
	for index, test := range tests {
		lists := []*ListNode{}
		for i, arr := range test.input {
			lists = append(lists, linked_list.InitList(arr))
			t.Logf("index: %v, input%d: %v", index, i, lists[i].ToString())
		}
		output := mergeKLists(lists)
		t.Logf("index: %v, output %v", index, output.ToString())
		assert.True(linked_list.InitList(test.output).IsEqual(output))
	}
}
