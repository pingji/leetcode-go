package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}
	for index, test := range tests {
		t.Logf("index: %v, nums: %v", index, test.nums)
		output := removeDuplicates(test.nums)
		t.Logf("index: %v, output %v", index, output)
		assert.Equal(test.output, output)
	}
}
