package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		nums   []int
		val    int
		output int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for index, test := range tests {
		t.Logf("index: %v, nums: %v, val %v", index, test.nums, test.val)
		output := removeElement(test.nums, test.val)
		t.Logf("index: %v, output %v", index, output)
		assert.Equal(test.output, output)
	}
}
