package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 1, 0}, []int{0, 0, 1}},
	}
	for index, test := range tests {
		t.Logf("index: %v, input: %v", index, test.input)
		sortColors(test.input)
		t.Logf("index: %v, output %v", index, test.input)
		assert.Equal(test.output, test.input)
	}
}
